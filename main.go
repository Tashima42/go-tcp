package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tashima42/go-tcp/client"
	"github.com/tashima42/go-tcp/server"
)

const (
	address  = "127.0.0.1:12345"
	protocol = "tcp"
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
		if err := server.Serve(address, protocol); err != nil {
			panic(err)
		}
	// caso seja client, inivia o client e envia a mensagem
	case "client":
		// verifica se tem os argumentos necessarios para rodar o client
		if len(os.Args) < 4 {
			printUsage()
			break
		}
		numberArg := os.Args[2]
		number, err := strconv.ParseFloat(numberArg, 64)
		name := os.Args[3]
		if err != nil || number <= 0 {
			log.Fatalf("falha ao converter argumento para numero, use um numero maior do que 0: %s", err.Error())
		}
		if err := client.SendMessage(address, protocol, number, name); err != nil {
			panic(err)
		}
	// caso seja qualquer outra coisa, mostre a mensagem de ajuda de uso
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Printf("usage: \n  %s server\n  %s client 5 pedro\n", os.Args[0], os.Args[0])
}
