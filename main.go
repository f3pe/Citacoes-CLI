package main

import (
	"log"
	"os"
	"citacoes-CLI/app"
)

func main()  {
	citation, err := app.NewCitation()
	if err != nil{
		log.Fatal(err)
	}

	if err := citation.Run(os.Args); err != nil{
		log.Fatal(err)
	}
}