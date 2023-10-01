package client

import (
	"fmt"
	"net"
)

func SendMessage(address, protocol string, number float64) error {
	// inicia uma conexao TCP
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return err
	}

	// envia a mensagem pela conexao iniciada
	if _, err := conn.Write([]byte(fmt.Sprintf("%f\n", number))); err != nil {
		return err
	}

	// fecha a conexao
	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
