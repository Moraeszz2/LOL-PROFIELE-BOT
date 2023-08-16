package lol

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SendRequest(nickname string) (getord interface{}) {

	envPath := filepath.Join("..", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Println(".env não foi possivel carregar", err)
		return
	}
	TOKEN := os.Getenv("TOKEN_LOL")
	URL := os.Getenv("URL_LOL")
	REGION := os.Getenv("REGION_LOL")

	if URL == "" || REGION == "" || TOKEN == "" {
		fmt.Println("URL/REGIAO/TOKEN do lol não foram fornecidos. Preencha os dados corretamente!")
		return
	}

	URL = URL + "summoner/v4/summoners/by-name/" + nickname

	req, err := http.NewRequest("GET", URL, nil)
	req.Header.Add("X-Riot-Token", TOKEN)

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	corpoResposta, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	// Verifica o código de status da resposta
	if resp.StatusCode == 200 {
		pl := &profileStruct{}
		json.Unmarshal([]byte(corpoResposta), pl)
		getord = map[string]interface{}{
			"name":          pl.Name,
			"summonerLevel": pl.SummonerLevel,
		}
		return getord
	} else {
		fmt.Printf("Erro: Código de status %d\n", resp.StatusCode)
	}

	return
}
