package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

type UsdQuote struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
	Error      string `json:"error"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("Erro do client ao fazer requisição: %s", err.Error()))
	}
	defer res.Body.Close()

	var quote UsdQuote
	if err := json.NewDecoder(res.Body).Decode(&quote); err != nil {
		panic(err)
	}

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Erro do server ao fazer requisição: status code %d, erro: %s", res.StatusCode, quote.Error))
	}

	writeInFile(quote.Bid)

	os.Stdout.WriteString("Dólar: " + quote.Bid)
}

func writeInFile(bid string) error {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	t := template.Must(template.New("cotacao").Parse("Dólar: {{.}}"))
	err = t.Execute(file, bid)

	return err
}
