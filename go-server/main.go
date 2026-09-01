package main

import (
	"fmt"
	"net"
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

		fmt.Printf("Recebido %d byte: %s", n, string(buf[:n]))

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Erro ao escrever para o cliente: ", err)
		}
	}
}
