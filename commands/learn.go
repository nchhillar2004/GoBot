package commands

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleLearn(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
    if len(args) == 0 {
        sendLearnHelp(s, m.ChannelID)
        return
    }

    subCommand := strings.ToLower(args[0])
    args = args[1:]

    switch subCommand{
    case "java":
        handleJavaLearn(s, m.ChannelID)
    case "cpp":
        handleCppLearn(s, m.ChannelID)
    case "mern":
        handleMernLearn(s, m.ChannelID)
    case "go":
        handleGoLearn(s, m.ChannelID)
    }
}

func handleJavaLearn(s *discordgo.Session, id string) {

}
func handleCppLearn(s *discordgo.Session, id string) {

}
func handleMernLearn(s *discordgo.Session, id string) {

}
func handleGoLearn(s *discordgo.Session, id string) {

}

func sendLearnHelp(s *discordgo.Session, id string) {
    helpMessage := `**Available resources:**
    !learn <skill> - Get the best learning resources for every skill, programming language
    Available commands:
    !learn c
    !learn cpp
    !learn gp
    !learn java
    !learn mern
    !learn go`

    _, err := s.ChannelMessageSendEmbed(id, &discordgo.MessageEmbed{
        Title:       "GoBot Learn Help",
        Description: helpMessage,
        Color:       000,
    })
    if err != nil {
        log.Printf("Error sending learn message: %v", err)
    }
}

