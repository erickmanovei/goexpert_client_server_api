# Desafio Client-Server-API

Este repositório se trata do desafio **Client-Server-API**, da Pós Graduação **Go Expert / FullCycle**. Nele estão contidos 2 projetos: um é o **server**, que se trata de um servidor HTTP que tem como objetivo informar a cotação do dólar do dia e persistir essa informação no SQLite, e o outro é o **client**, que tem como objetivo consumir o server, coletar a cotação e gravar em um arquivo local.

Ambos os projetos têm como finalidade o exercício e a prática de técnicas como context, HTTP Server, manipulação de arquivos, structs, ponteiros, manipulação de JSON, templates dentre outros.

Nestes projetos, não foram aplicadas algumas boas práticas, como a arquitetura REST, com o objetivo de manter o escopo conforme o desafio.

## Escopo do Desafio

Você precisará nos entregar dois sistemas em Go:
- `client.go`
- `server.go`

Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Como Executar os Projetos

### Server

1. Navegue até o diretório do projeto **server**.
2. Execute o comando `go run main.go` para iniciar o servidor HTTP.

### Client

1. Navegue até o diretório do projeto **client**.
2. Execute o comando `go run main.go` para solicitar a cotação do dólar e salvar a informação no arquivo local.

### Extra (fora do escopo)

1. Foi criado o endpoint `GET /cotacao/list` com o objetivo de testar a persistência de dados no SQLite;
2. Foi adicionado o arquivo `/server/.vscode/launch.json` para possibilitar debug do `server`. Para executar o debug, basta ter a extensão oficial do Go no VS Code e pressionar F5, estando com o arquivo `/server/main.go` aberto.

### Observações Importantes

Existe uma grande probabilidade de ocorrerem erros de TimeOut, principalmente no ato da persistência de dados no SQLite, já que o tempo pedido de time out no escopo para esse ponto foi de apenas 10ms.
Dessa forma, lebre-se de aumentar esses tempos no arquivo `server/controllers/quote_controller.go`, linhas 16 e 18, e no arquivo `client/main.go`, linha 29, para que os testes ocorram sem excessões.

