package main

import (
	"fmt"
	"net"
	"slices"
	"strconv"
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
		command, err := validateOperation(string(buf[:n]))
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			break
		}

		operation, values, err := parseOperation(command)
		if err != nil {
			fmt.Printf("Erro: %v\n", err)
			break
		}
		fmt.Printf("Operação %s sobre %v\n", operation, values)

		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Erro ao escrever para o cliente: ", err)
		}
	}
}

func validateOperation(input string) ([]string, error) {
	command := strings.Split(input, " ")
	fmt.Printf("%q\n", command)
	fmt.Printf("%d\n", len(command))

	if len(command) < 3 {
		return command, fmt.Errorf("comando de input inválido: %q", command)
	}

	validOperations := []string{"SOMA", "SUB", "MUL", "DIV"}
	operation := strings.ToUpper(command[0])
	if !slices.Contains(validOperations, operation) {
		return command, fmt.Errorf("operação inválida: %q", operation)
	}

	command[0] = operation

	return command, nil
}

func parseOperation(command []string) (string, []float64, error) {
	operation := command[0]

	values := make([]float64, 0, len(command)-1)
	for _, raw := range command[1:] {
		value, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return operation, nil, fmt.Errorf("operando inválido %q: %w", raw, err)
		}
		values = append(values, value)
	}

	return operation, values, nil
}

func sum(values []float64) (float64) {
	sum := values[0] + values[1]
	return sum
}

func div(values []float64) (float64, error) {
	if values[1] == 0 {
		return 0, fmt.Errorf("Divisão por zero!")
	}

	div := values[0] / values[1]
	return div, nil
}

func switchOperation(operation string, values []float64) {

	switch operation {
		case "SOMA":
			resultado := sum(values)
			fmt.Printf("Resultado da soma => %f\n", resultado)
			return
		case "SUB":
			// TO DO
		case "MUL":
			// TO DO
		case "DIV":
			resultado, err := div(values)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Resultado da divisão => %f\n", resultado)
	}
}
