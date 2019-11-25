package web开发

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

// 中间件
func timeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)

		fmt.Println(timeElapsed)
	})
}

func TestMiddleware(t *testing.T) {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
