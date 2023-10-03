package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func SendMessage(address, protocol string, number float64, name string) error {
	// inicia uma conexao TCP
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return err
	}
	// envia a mensagem pela conexao iniciada
	if _, err := conn.Write([]byte(fmt.Sprintf("%f/&/%s\n", number, name))); err != nil {
		return err
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		log.Print("resultado do calculo: " + scanner.Text())
	}

	// fecha a conexao
	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
