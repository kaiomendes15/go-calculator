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
			responderErro(conn, err)
			continue
		}

		operation, values, err := parseOperation(command)
		if err != nil {
			responderErro(conn, err)
			continue
		}
		fmt.Printf("Operação %s sobre %v\n", operation, values)
		resultado, err := switchOperation(operation, values)
		if err != nil {
			responderErro(conn, err)
			continue
		}

		resposta := strconv.FormatFloat(resultado, 'f', -1, 64)
		_, err = conn.Write([]byte(resposta))
		if err != nil {
			fmt.Println("Erro ao escrever para o cliente: ", err)
		}

	}
}

func validateOperation(input string) ([]string, error) {
	command := strings.Fields(input)
	fmt.Printf("%q\n", command)
	fmt.Printf("%d\n", len(command))

	if len(command) == 0 {
		return command, fmt.Errorf("comando vazio")
	}

	validOperations := []string{"SOMA", "SUB", "MUL", "DIV"}
	operation := strings.ToUpper(command[0])
	if !slices.Contains(validOperations, operation) {
		return command, fmt.Errorf("comando desconhecido: %q", operation)
	}

	if len(command) != 3 {
		return command, fmt.Errorf("formato invalido (use: OPERACAO NUM1 NUM2)")
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

func sum(values []float64) float64 {
	sum := values[0] + values[1]
	return sum
}

func div(values []float64) (float64, error) {
	if values[1] == 0 {
		return 0, fmt.Errorf("divisão por zero!")
	}

	div := values[0] / values[1]
	return div, nil
}

func mul(values []float64) float64 {
	mul := values[0] * values[1]
	return mul
}

func sub(values []float64) float64 {
	sub := values[0] - values[1]
	return sub
}

func switchOperation(operation string, values []float64) (float64, error) {

	resultado := 0.0

	switch operation {
	case "SOMA":
		resultado = sum(values)
		fmt.Printf("Resultado da soma => %f\n", resultado)
		return resultado, nil

	case "SUB":
		resultado = sub(values)
		fmt.Printf("Resultado da subtração => %f\n", resultado)
		return resultado, nil

	case "MUL":
		resultado = mul(values)
		fmt.Printf("Resultado da multiplicação => %f\n", resultado)
		return resultado, nil

	case "DIV":
		resultado, err := div(values)
		if err != nil {
			return resultado, err
		}

		fmt.Printf("Resultado da divisão => %f\n", resultado)
	}

	return resultado, nil
}

func responderErro(conn net.Conn, erro error) {
	fmt.Printf("Erro: %v\n", erro)
	if _, err := conn.Write([]byte("ERRO: " + erro.Error())); err != nil {
		fmt.Println("Erro ao escrever para o cliente: ", err)
	}
}
