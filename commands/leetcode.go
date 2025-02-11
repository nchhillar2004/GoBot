package commands

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/nchhillar2004/gobot/utils"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func HandleLeetCode(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
    if len(args) == 0 {
		sendLeetCodeHelp(s, m.ChannelID)
		return
	}

	subCommand := strings.ToLower(args[0])
	args = args[1:]

	switch subCommand {
	case "random":
		handleRandomQuestions(s, m, args)
	default:
		sendLeetCodeHelp(s, m.ChannelID)
	}
}

func handleRandomQuestions(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
    if len(args) == 1 {
		handleRandomDifficultyQuestions(s, m.ChannelID, args[0])
		return
	}

    question, err := utils.GetRandomLeetCodeQuestion()
    if err != nil {
        log.Fatalf("Error fetching random leetcode question: %v", err)
        return
    }

    url := "https://leetcode.com/problems/" + question.TitleSlug

    author := discordgo.MessageEmbedAuthor{
        Name: "Leetcode random",
        URL: url,
    }

    embed := discordgo.MessageEmbed{
        Title: question.Title + " (" + question.Difficulty + ")",
        Author: &author,
        URL: url,
    }

    s.ChannelMessageSendEmbed(m.ChannelID, &embed)
}

func handleRandomDifficultyQuestions(s *discordgo.Session, id string, arg string) {
    question, err := utils.GetRandomLeetCodeQuestion(arg)
    if err != nil {
        if err == utils.ErrInvalidDifficulty {
            s.ChannelMessageSend(id, "Invalid difficulty, Select from 'easy' OR 'medium' OR 'hard'")
        }
        return
    }

    url := "https://leetcode.com/problems/" + question.TitleSlug

    author := discordgo.MessageEmbedAuthor{
        Name: "Leetcode " + arg,
        URL: url,
    }

    embed := discordgo.MessageEmbed{
        Title: question.Title + " (" + arg + ")",
        Author: &author,
        URL: url,
    }

    s.ChannelMessageSendEmbed(id, &embed)
}

func sendLeetCodeHelp(s *discordgo.Session, id string) {
    helpMessage := `**Leetcode Available Commands:**
    !leetcode <company-name>
    !leetcode random
    !leetcode random <easy | medium | hard>
    !help - Show all commands`

    _, err := s.ChannelMessageSendEmbed(id, &discordgo.MessageEmbed{
        Title:       "GoBot Leetcode Help",
        Description: helpMessage,
        Color:       808080,
    })
    if err != nil {
        log.Printf("Error sending leetcode help message: %v", err)
    }
}
