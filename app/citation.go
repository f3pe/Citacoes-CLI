package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

type citation struct{
	Quote 		string `json:"q"`
	Author 		string `json:"a"`
	Img			string `json:"i"`
	Characters 	string `json:"c"`
	HTML 		string `json:"h"`
}

func NewCitation() (*cli.App ,error)  {
	app := &cli.App{
		Name: "App de citações",
		Usage: "Gera uma sitação aleatoria com base na API zenquotes",
		Action: generateCitation,
	}
	return app, nil
}


func generateCitation(cliCtx *cli.Context) error {
	response, err := http.Get("https://zenquotes.io/api/random")
	if err != nil{
		log.Fatal("não foi possivel obter a citação error:", err)
	}
	defer response.Body.Close()


	body, err := io.ReadAll(response.Body)
	if err != nil{
		log.Fatal("Não foi possivel ler a resposta do servidor error:", err)
	}

	var cit []citation


	err = json.Unmarshal(body, &cit)
	if err != nil{
		log.Fatal("Error ao analisar JSON errro:", err)
	}

	fmt.Printf("Citation: \"%s\" - %s\n", cit[0].Quote, cit[0].Author)

	return nil
}