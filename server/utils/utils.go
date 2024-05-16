package utils

import (
	"encoding/json"
	"net/http"
)

func HttpResponse(w *http.ResponseWriter, status int, response any) {
	header := (*w).Header()
	header.Set("Content-Type", "application/json")
	(*w).WriteHeader(status)
	json.NewEncoder(*w).Encode(response)
}
