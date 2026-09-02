# 🧮 Go Calculator — Servidor de Cálculo com Sockets TCP

Atividade prática de **Redes de Computadores**: uma aplicação cliente-servidor
escrita em Go que funciona como uma **calculadora remota**.

O cliente envia comandos digitados pelo usuário; o servidor recebe o comando pela
conexão TCP, faz a conta e devolve o resultado pela mesma conexão. A conexão fica
aberta para vários comandos e cada cliente é atendido em sua própria goroutine.

## 🛠️ Tecnologias

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![TCP Sockets](https://img.shields.io/badge/TCP-Sockets-4479A1?style=for-the-badge&logo=cloudflare&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)

- **Go (`net`)** — `net.Listen` no servidor e `net.Dial` no cliente
- **Goroutines** — um cliente por goroutine, vários clientes ao mesmo tempo
- **`bufio.Scanner`** — leitura da entrada do usuário no cliente
- **`strings.Fields` / `strconv`** — parsing e conversão dos operandos

## 📡 Protocolo

Mensagens de texto no formato:

```
OPERACAO OPERANDO1 OPERANDO2
```

| Comando | Descrição                     | Exemplo     |
|---------|-------------------------------|-------------|
| `SOMA`  | Soma dois números             | `SOMA 10 20`|
| `SUB`   | Subtrai o segundo do primeiro | `SUB 50 8`  |
| `MUL`   | Multiplica dois números       | `MUL 6 7`   |
| `DIV`   | Divide o primeiro pelo segundo| `DIV 20 4`  |
| `SAIR`  | Encerra a conexão (no cliente)| `SAIR`      |

### Tratamento de erros

- **Divisão por zero** → `ERRO: divisão por zero!`
- **Comando desconhecido** → `ERRO: comando desconhecido: "XPTO"`
- **Formato inválido** (nº de argumentos errado ou operando não numérico) →
  `ERRO: formato invalido (use: OPERACAO NUM1 NUM2)`

## 📁 Estrutura

```
go-calculator/
├── go-server/main.go   # servidor TCP (porta 8090)
└── client/client.go    # cliente CLI interativo
```

## ▶️ Como executar

Pré-requisito: **Go 1.22+** instalado.

### 1. Suba o servidor

```bash
cd go-server
go run main.go
```

Saída esperada:

```
Echo server rodando na porta 8090...
```

### 2. Em outro terminal, rode o cliente

```bash
cd client
go run client.go
```

### 3. Use a calculadora

```
calc> SOMA 10 20
Resultado: 30
calc> DIV 10 0
Resultado: ERRO: divisão por zero!
calc> HELLO
Resultado: ERRO: comando desconhecido: "HELLO"
calc> MUL 3
Resultado: ERRO: formato invalido (use: OPERACAO NUM1 NUM2)
calc> SAIR
```

O comando `SAIR` fecha a conexão e encerra o cliente. Você pode abrir vários
clientes ao mesmo tempo — o servidor atende todos simultaneamente.

## 👥 Autores

Trabalho em dupla — Redes de Computadores.
