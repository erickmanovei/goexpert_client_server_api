package services

import (
	"context"
	"database/sql"

	"github.com/erickmanovei/client_server_api_server/structs"
	_ "github.com/mattn/go-sqlite3"
)

func GetAllQuotes(ctx *context.Context) ([]structs.SimpleUsdQuote, error) {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(*ctx, "SELECT code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date FROM cotacoes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []structs.SimpleUsdQuote
	for rows.Next() {
		var quote structs.SimpleUsdQuote
		err := rows.Scan(&quote.Code, &quote.Codein, &quote.Name, &quote.High, &quote.Low, &quote.VarBid, &quote.PctChange, &quote.Bid, &quote.Ask, &quote.Timestamp, &quote.CreateDate)
		if err != nil {
			return nil, err
		}
		quotes = append(quotes, quote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}
