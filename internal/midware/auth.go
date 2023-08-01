package midware

import (
	"fmt"
	"net/http"
)

func MockupAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		fmt.Println("user authenticated!")
		next.ServeHTTP(w, r)
	})
}