package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Printf("calc> ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Erro ao ler:", err)
			}
			break
		}

		text := scanner.Text()

		if text == "SAIR" {
			break
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Erro ao enviar dados:", err)
			return
		}

		fmt.Printf("Enviado: %s\n", text)

		buf := make([]byte, 1024)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Erro ao ler resposta do servidor:", err)
			return
		}

		fmt.Printf("Resultado: %s\n", string(buf[:n]))

	}
}
