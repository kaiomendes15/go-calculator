package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	message := "Olá, servidor Echo com Buffer em Go!\n"

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Erro ao enviar dados:", err)
		return
	}

	fmt.Printf("Enviado: %s", message)

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Erro ao ler resposta do servidor:", err)
		return
	}

	fmt.Printf("Echo recebido do servidor (%d bytes): %s", n, string(buf[:n]))

}
