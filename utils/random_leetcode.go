package utils

import (
	"encoding/json"
	"fmt"
    "errors"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	ErrInvalidDifficulty = errors.New("Invalid difficulty: please provide 'easy', 'medium', or 'hard'")
)

type Question struct {
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
}

var (
	cacheMutex     sync.RWMutex
	questionsCache struct {
		all    []Question
		easy   []Question
		medium []Question
		hard   []Question
		expiry time.Time
	}
	cacheTTL = 30 * time.Minute
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetRandomLeetCodeQuestion(difficulty ...string) (Question, error) {
	if err := refreshCacheIfNeeded(); err != nil {
		return Question{}, err
	}

	cacheMutex.RLock()
	defer cacheMutex.RUnlock()

	var targetPool []Question
	if len(difficulty) > 0 {
		diff := strings.ToLower(difficulty[0])
		switch diff {
		case "easy":
			targetPool = questionsCache.easy
		case "medium":
			targetPool = questionsCache.medium
		case "hard":
			targetPool = questionsCache.hard
        default:
            return Question{}, ErrInvalidDifficulty
		}
	} else {
		targetPool = questionsCache.all
	}

	if len(targetPool) == 0 {
		return Question{}, fmt.Errorf("no questions available")
	}

	return targetPool[rand.Intn(len(targetPool))], nil
}

func refreshCacheIfNeeded() error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if time.Now().Before(questionsCache.expiry) {
		return nil
	}

	resp, err := http.Get("https://leetcode.com/api/problems/all/")
	if err != nil {
		return fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	var data struct {
		Problems []struct {
			Stat       struct {
				Title     string `json:"question__title"`
				TitleSlug string `json:"question__title_slug"`
			} `json:"stat"`
			Difficulty struct {
				Level int `json:"level"`
			} `json:"difficulty"`
		} `json:"stat_status_pairs"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("JSON decode failed: %w", err)
	}

	newCache := struct {
		all    []Question
		easy   []Question
		medium []Question
		hard   []Question
	}{}

	for _, p := range data.Problems {
		var diff string
		switch p.Difficulty.Level {
		case 1:
			diff = "easy"
		case 2:
			diff = "medium"
		case 3:
			diff = "hard"
		default:
			continue
		}

		q := Question{
			Title:      p.Stat.Title,
			TitleSlug:  p.Stat.TitleSlug,
			Difficulty: diff,
		}

		newCache.all = append(newCache.all, q)
		switch diff {
		case "easy":
			newCache.easy = append(newCache.easy, q)
		case "medium":
			newCache.medium = append(newCache.medium, q)
		case "hard":
			newCache.hard = append(newCache.hard, q)
		}
	}

	questionsCache.all = newCache.all
	questionsCache.easy = newCache.easy
	questionsCache.medium = newCache.medium
	questionsCache.hard = newCache.hard
	questionsCache.expiry = time.Now().Add(cacheTTL)

	return nil
}
