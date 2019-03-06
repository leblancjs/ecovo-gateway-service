package handler

import "net/http"

// HelloWorld says hello to the world.
func HelloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!\n"))
	})
}
