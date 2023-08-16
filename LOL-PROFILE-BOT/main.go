package main

import (
	help "LOL-PROFILE-BOT/LOL-PROFILE-BOT/Help"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println(".env não foi possivel carregar", err)
		return
	}
	TOKEN := os.Getenv("TOKEN_DISCORD")

	if TOKEN == "" {
		fmt.Println("Token do bot não fornecido. Preencha os dados corretamente!")
		return
	}

	DC, err := discordgo.New("Bot " + TOKEN)

	if err != nil {
		fmt.Println("Erro ao criar a sessão do bot: ", err)
		return
	}

	DC.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID || m.Author.Bot {
			return
		}

		switch m.Content {
		case "Bot":
			resp := "Olá " + m.Author.Username + ", como posso ajudar?"
			s.ChannelMessageSend(m.ChannelID, resp)

			userResp, err := help.WaitForUserResponse(s, m.ChannelID, m.Author.ID)
			if err != nil {
				fmt.Println("Erro ao aguardar resposta do usuário:", err)
				return
			}

			if userResp == "comandos" {
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
				_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
				if err != nil {
					fmt.Println("Error sending embed:", err)
				}

			} else if userResp == " o tempo limite foi atingido para resposta!" {
				s.ChannelMessageSend(m.ChannelID, m.Author.Username+userResp)
			} else {
				err := "Comando não encontrado"
				s.ChannelMessageSend(m.ChannelID, err)
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
