package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/erickmanovei/client_server_api_server/services"
	"github.com/erickmanovei/client_server_api_server/structs"
	"github.com/erickmanovei/client_server_api_server/utils"
)

func ShowQuote(w http.ResponseWriter, r *http.Request) {
	ctxUsdQuote, cancelUsdQuote := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancelUsdQuote()
	ctxPersistQuote, cancelPersistQuote := context.WithTimeout(context.Background(), time.Millisecond*10)
	defer cancelPersistQuote()

	usdQuote := structs.UsdQuote{}
	errorResponse := structs.ErrorResponse{}

	err := services.GetUsdQuote(&ctxUsdQuote, &usdQuote)
	if err != nil {
		log.Println(err.Error())
		errorResponse.Error = err.Error()
		utils.HttpResponse(&w, http.StatusInternalServerError, errorResponse)
		return
	}

	err = services.PersistUsdQuote(&ctxPersistQuote, &usdQuote)
	if err != nil {
		log.Println(err.Error())
		errorResponse.Error = err.Error()
		utils.HttpResponse(&w, http.StatusInternalServerError, errorResponse)
		return
	}

	utils.HttpResponse(&w, http.StatusOK, usdQuote.USDBRL)
}

func ListQuotes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	quotes, err := services.GetAllQuotes(&ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
