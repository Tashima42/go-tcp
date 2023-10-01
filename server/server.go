package server

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

const (
	multiply = 100.00 // adryann
	divide   = 101.00 // eloz
)

func Serve(address, protocol string) error {
	log.Println("iniciando servidor...")
	// net.Listen anuncia na rede e espera por conexoes, nesse caso TCP
	listen, err := net.Listen(protocol, address)
	if err != nil {
		return err
	}
	log.Println("aceitando conexoes " + protocol + " no endereco: " + address)
	// em go, a keyword defer e usada pra executar uma acao logo antes de
	// retornar a funcao, nesse caso logo antes de retornar a funcao, o
	// servidor sera fechado
	defer listen.Close()
	// escutando continuamente por novas conexoes
	for {
		// listen.Accept, quando uma conexao e iniciada, aceita e retorna ela
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		// a mensagem do usuario e colocada em um buffer e depois lida ate o delimitador de nova linha
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return err
		}
		// o delimitador de nova linha e removido para nao atrapalhar na conversao para numero
		data = strings.Split(data, "\n")[0]
		log.Println("numero recebido: ", data)
		number, err := strconv.ParseFloat(data, 64)
		if err != nil {
			return err
		}
		finalNumber := (number * multiply) / divide
		log.Printf("(%.2f * %.2f) / %.2f: %.2f", number, multiply, divide, finalNumber)
		// assim que a mensagem e lida, a conexao com o cliente e fechada (FIN)
		if err := conn.Close(); err != nil {
			return err
		}
	}
}
