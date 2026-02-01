package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONError(w http.ResponseWriter, message string, httpStatus int) {
    errorResponse := map[string]string{
        "error": message,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(httpStatus)
    json.NewEncoder(w).Encode(errorResponse)
}
