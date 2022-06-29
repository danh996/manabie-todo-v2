package transport

import (
	"encoding/json"
	"net/http"
)

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	data := map[string]interface{}{
		"data": object,
	}
	json.NewEncoder(writer).Encode(data)
}
