package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type Configuration struct {
	DiscordToken string
	Prefix       string
	LolToken     string
	UrlLol       string
	RegiaoLol    string
}

func LoadConfig() (Configuration, error) {
	envPath := filepath.Join("..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		return Configuration{}, err
	}

	config := Configuration{
		DiscordToken: os.Getenv("TOKEN_DISCORD"),
		Prefix:       os.Getenv("PREFIX"),
		LolToken:     os.Getenv("TOKEN_LOL"),
		UrlLol:       os.Getenv("URL_LOL"),
		RegiaoLol:    os.Getenv("REGION_LOL"),
	}

	return config, nil
}

func ConnectDiscord(token string) (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar a sessão do bot: %w", err)
	}

	return dg, nil
}
