package main

import (
	"fmt"
	"os"

	"github.com/tashima42/go-tcp/client"
	"github.com/tashima42/go-tcp/server"
)

func main() {
	// modo pode ser server ou client
	var mode string
	if len(os.Args) < 2 {
		mode = "help"
	} else {
		mode = os.Args[1]
	}
	switch mode {
	// caso seja server, inicia o servidor
	case "server":
		if err := server.Serve(); err != nil {
			panic(err)
		}
	// caso seja client, inivia o client e envia a mensagem
	case "client":
		// verifica se tem os argumentos necessarios para rodar o client
		if len(os.Args) < 3 {
			printUsage()
			break
		}
		message := os.Args[2]
		if err := client.SendMessage(message); err != nil {
			panic(err)
		}
	// caso seja qualquer outra coisa, mostre a mensagem de ajuda de uso
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Printf("usage: \n  %s server\n  %s client teste\n", os.Args[0], os.Args[0])
}
