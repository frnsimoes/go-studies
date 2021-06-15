package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("ponto de partida")

	app := app.Gerar()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
