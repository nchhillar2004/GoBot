package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/nchhillar2004/gobot/utils"
)

func HandleInterview(s *discordgo.Session, m *discordgo.MessageCreate) {
    questions := utils.GetCompanyQuestions()

    _, err := s.ChannelMessageSend(m.ChannelID, questions)
    if err != nil {
        log.Fatalf("Error sending questions: %v", err)
    }
}
