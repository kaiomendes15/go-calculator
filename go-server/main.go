package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Echo server rodando na porta 8090...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar uma conexão: ", err)
			continue
		}

		//           coroutine
		// Criar uma goroutine aqui
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Cliente conectado %s\n", conn.RemoteAddr())
	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Cliente desconectado ou erro de leitura: %v\n", err)
			break
		}

		fmt.Printf("Recebido %d byte: %s\n", n, string(buf[:n]))
		operation, err := validateOperation(string(buf[:n]))
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			break
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Erro ao escrever para o cliente: ", err)
		}
	}
}

func validateOperation(input string) ([]string, error) {
	operation := strings.Split(input, " ")
	fmt.Printf("%q\n", operation)
	fmt.Printf("%d\n", len(operation))
	
	if len(operation) < 3 {
		return operation, fmt.Errorf("Comando de input inválido: %q\n", operation)
	} 

	return operation, nil
}
