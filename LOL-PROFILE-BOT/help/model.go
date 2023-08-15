package help

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func WaitForUserResponse(s *discordgo.Session, channelID, userID string) (string, error) {
	responseCh := make(chan string)
	timeoutCh := time.After(15 * time.Second)

	s.AddHandlerOnce(func(innerS *discordgo.Session, innerM *discordgo.MessageCreate) {
		if innerM.Author.ID == userID && innerM.ChannelID == channelID {
			responseCh <- innerM.Content
		}
	})

	select {
	case response := <-responseCh:
		return response, nil
	case <-timeoutCh:
		return " o tempo limite foi atingido para resposta!", nil
	}
}
