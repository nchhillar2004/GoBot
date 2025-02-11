package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
    if len(args) == 0 {
        sendHelp(s, m.ChannelID)
        return
    }

    subCommand := strings.ToLower(args[0])
    args = args[1:]

    switch subCommand{
    case "leetcode":
        sendLeetCodeHelp(s, m.ChannelID)
    case "learn":
        sendLearnHelp(s, m.ChannelID)
    }
}

func sendHelp(s *discordgo.Session, id string) {
    helpMessage := "```\nAll Available Commands:\n" +
    "!learn <skill>            - Get learning resources for skill\n" +
    "!interview                - Get top interview questions (last 6 months)\n" +
    "!leetcode <company-name>  - Get top 20 questions from company\n" +
    "!leetcode random          - Get a random LeetCode question\n" +
    "!leetcode random <diff>   - Get a random question by difficulty\n" +
    "!cp                       - Get competitive programming problems\n" +
    "!jobs                     - Fetch latest job listings\n" +
    "!help <command>           - Get help for a specific command, Ex. !help leetcode\n" +
    "!help                     - Show this help message\n```"

    _, err := s.ChannelMessageSend(id, helpMessage)
    if err != nil {
        log.Printf("Error sending help message: %v", err)
    }
}
