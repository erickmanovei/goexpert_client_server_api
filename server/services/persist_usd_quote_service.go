package services

import (
	"context"
	"database/sql"
	"log"

	"github.com/erickmanovei/client_server_api_server/structs"
	_ "github.com/mattn/go-sqlite3"
)

func PersistUsdQuote(ctx *context.Context, usdQuote *structs.UsdQuote) error {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
	);
	`
	_, err = db.ExecContext(*ctx, sqlStmt)
	if err != nil {
		return err
	}

	insertSQL := `INSERT INTO cotacoes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(usdQuote.USDBRL.Code, usdQuote.USDBRL.Codein, usdQuote.USDBRL.Name, usdQuote.USDBRL.High, usdQuote.USDBRL.Low, usdQuote.USDBRL.VarBid, usdQuote.USDBRL.PctChange, usdQuote.USDBRL.Bid, usdQuote.USDBRL.Ask, usdQuote.USDBRL.Timestamp, usdQuote.USDBRL.CreateDate)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
