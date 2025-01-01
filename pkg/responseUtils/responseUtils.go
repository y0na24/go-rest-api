package responseUtils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}
