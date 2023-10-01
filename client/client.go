package client

import (
	"net"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1:12345"
)

func SendMessage(message string) error {
	// inicia uma conexao TCP
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return err
	}

	// envia a mensagem pela conexao iniciada
	if _, err := conn.Write([]byte(message)); err != nil {
		return err
	}

	// fecha a conexao
	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
