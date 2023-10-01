package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	address  = "127.0.0.1:12345"
	protocol = "tcp"
)

func Serve() error {
	log.Println("iniciando servidor...")
	// net.Listen anuncia na rede e espera por conexoes, nesse caso TCP
	listen, err := net.Listen(protocol, address)
	if err != nil {
		return err
	}
	log.Println("aceitando conexoes " + protocol + " no endereco: " + address)
	// listen.Accept, quando uma conexao e iniciada, aceita e retorna ela
	conn, err := listen.Accept()
	if err != nil {
		return err
	}
	// os dados chegam em bytes e depois podem ser convetidas para strings
	var data []byte
	if _, err := conn.Read(data); err != nil {
		return err
	}
	// a mensgem aqui e convertida para uma string e printada no terminal
	log.Println(string(data))
	// assim que a mensagem e lida, a conexao com o cliente e fechada (FIN)
	if err := conn.Close(); err != nil {
		return err
	}

	// o programa nao vai avancar, ate que o ususario aperte ctrl + c no terminal
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	log.Println("Pressione ctrl+c para fechar o servidor")
	<-done
	log.Println("fechando servidor...")
	// quando o usuario pressionar ctrl + c, o servidor vai ser encerrado
	return listen.Close()
}
