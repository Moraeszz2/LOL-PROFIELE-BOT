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

func a() {

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
		case "Bot", "bot":
			resp := "Olá " + m.Author.Username + ", como posso ajudar?"
			s.ChannelMessageSend(m.ChannelID, resp)

			userResp, err := help.WaitForUserResponse(s, m.ChannelID, m.Author.ID)
			if err != nil {
				fmt.Println("Erro ao aguardar resposta do usuário:", err)
				return
			}

			if userResp == "comandos" {
				// commandList := "1. comando do lol\n2. comando do bot\n3. sair"
				// embed := &discordgo.MessageEmbed{
				// 	Title:       "Comandos",
				// 	Description: commandList,
				// 	Color:       0x8905b5,
				// 	Thumbnail: &discordgo.MessageEmbedThumbnail{
				// 		URL:    m.Author.AvatarURL("256"), // Tamanho do ícone (16, 32, 64, 128, 256, 512)
				// 		Width:  256,
				// 		Height: 256,
				// 	},
				// }
				option1 := discordgo.SelectMenuOption{
					Label: "Option 1",
					Value: "option1",
				}

				option2 := discordgo.SelectMenuOption{
					Label: "Option 2",
					Value: "option2",
				}

				// Create the select menu
				selectMenu := discordgo.SelectMenu{
					CustomID:    "selectmenu1", // A unique identifier for the menu
					Placeholder: "Select an option",
					Options:     []discordgo.SelectMenuOption{option1, option2},
				}

				// Create an action row with the select menu
				component := discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{&selectMenu},
				}

				// Create a message with the action row and select menu
				message := discordgo.MessageSend{
					Content:    "Choose an option:",
					Components: []discordgo.MessageComponent{&component},
				}

				// Send the message with the select menu
				s.ChannelMessageSendComplex(m.ChannelID, &message)

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
