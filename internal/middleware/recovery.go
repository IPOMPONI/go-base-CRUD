package middleware

import (
	"log"
	"net/http"

	"booklib/internal/pkg/utils"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("[FATAL]: panic recovered %v", err)
				utils.SendJSONError(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
