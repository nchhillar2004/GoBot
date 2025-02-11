package handlers

import (
    "log"
    "strings"

    "github.com/bwmarrin/discordgo"
    "github.com/nchhillar2004/gobot/commands"
)

func InitHandlers(s *discordgo.Session) {
    s.AddHandler(messageCreate)
    s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
        log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
    })
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    if m.Author.ID == s.State.User.ID || m.Author.Bot {
        return
    }

    prefix := "!"

    if !strings.HasPrefix(m.Content, prefix) {
        return
    }

    args := strings.Fields(m.Content[len(prefix):])
    if len(args) < 1 {
        return
    }

    command := strings.ToLower(args[0])
    args = args[1:]

    switch command {
    case "leetcode":
        commands.HandleLeetCode(s, m, args)
    case "cp":
        commands.HandleCP(s, m, args)
    case "interview":
        commands.HandleInterview(s, m)
    case "jobs":
        commands.HandleJobs(s, m, args)
    case "learn":
        commands.HandleLearn(s, m, args)
    case "help":
        commands.HandleHelp(s, m, args)
    }
}
