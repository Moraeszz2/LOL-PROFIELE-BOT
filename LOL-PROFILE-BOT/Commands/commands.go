package commands

import (
	config "LOL-PROFILE-BOT/LOL-PROFILE-BOT/Config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Anything() *discordgo.MessageEmbed {
	configData, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Erro ao carregar configurações:", err)
		return nil
	}

	DC, err := config.ConnectDiscord(configData.DiscordToken)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Criar um canal para sinalizar que o handler foi concluído
	done := make(chan struct{})

	// Variável para armazenar o embed
	var embed *discordgo.MessageEmbed

	DC.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		defer close(done) // Sinaliza que o handler foi concluído

		commandList := "1. comando do lol\n2. comando do bot\n3. sair"

		embed = &discordgo.MessageEmbed{
			Title:       "Comandos",
			Description: commandList,
			Color:       0x8905b5,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL:    m.Author.AvatarURL("256"),
				Width:  256,
				Height: 256,
			},
		}
	})
	return embed

}
