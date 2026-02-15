package middleware

import (
	"log"
	"net/http"
	"strconv"

	"booklib/internal/pkg/utils"
)

func CheckBookIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))

		if err != nil {
			log.Printf("[ERROR]: %s %s - %s", r.Method, r.URL, err.Error())
			utils.SendJSONError(w, "'id' is invalid", http.StatusBadRequest)
			return
		}

		if id < 0 {
			log.Printf("[ERROR]: %s %s - 'id' is negative", r.Method, r.URL)
			utils.SendJSONError(w, "'id' is negative", http.StatusBadRequest)
			return
		}
		ctx := utils.SetBookIdInCtx(r.Context(), id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
