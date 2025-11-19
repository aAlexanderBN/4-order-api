package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "applicatin/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(data)
}
