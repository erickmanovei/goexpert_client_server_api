package routes

import "net/http"

func Routes(mux *http.ServeMux) {
	QuoteRoutes(mux)
}
