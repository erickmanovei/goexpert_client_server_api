package routes

import (
	"net/http"

	"github.com/erickmanovei/client_server_api_server/controllers"
)

func QuoteRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/cotacao", controllers.ShowQuote)       // não segui o padrão REST para manter conformidade com o pedido no escopo
	mux.HandleFunc("/cotacao/list", controllers.ListQuotes) // endpoint fora do escopo, mas criado para testar a persistência de dados
}
