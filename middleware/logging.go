package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("requests.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		logger := log.New(file, "", log.LstdFlags)
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		logger.Printf("%s %s %s\n", r.Method, r.RequestURI, duration)
	})
}
