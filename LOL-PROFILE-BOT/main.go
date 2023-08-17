package main

import (
	config "LOL-PROFILE-BOT/LOL-PROFILE-BOT/Config"
	help "LOL-PROFILE-BOT/LOL-PROFILE-BOT/Help"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	configData, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Erro ao carregar configurações:", err)
		return
	}

	DC, err := config.ConnectDiscord(configData.DiscordToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	DC.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if !strings.HasPrefix(m.Content, configData.Prefix) {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] == configData.Prefix+"bot" {
			respBot := "Olá " + m.Author.Username + ", como posso ajudar?"
			s.ChannelMessageSend(m.ChannelID, respBot)

			respUser, err := help.WaitForUserResponse(s, m.ChannelID, m.Author.ID)

			if err != nil {
				fmt.Println("Erro ao aguardar resposta do usuário:", err)
				return
			}

			if respUser == configData.Prefix+"comandos" {
				commandList := "1. comando do lol\n2. comando do bot\n3. sair"

				embed := &discordgo.MessageEmbed{
					Title:       "Comandos",
					Description: commandList,
					Color:       0x8905b5,
					Thumbnail: &discordgo.MessageEmbedThumbnail{
						URL:    m.Author.AvatarURL("256"), // Tamanho do ícone (16, 32, 64, 128, 256, 512)
						Width:  256,
						Height: 256,
					},
				}
				s.ChannelMessageSendEmbed(m.ChannelID, embed)

			}

		}
	})

	DC.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = DC.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer DC.Close()
	fmt.Println("Bot esta on!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

}
