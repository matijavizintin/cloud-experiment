package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", &indexHandler{})
	http.Handle("/healthcheck/", &healthCheckHandler{})

	http.ListenAndServe(":8080", nil)
}

type indexHandler struct{}

func (i indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(fmt.Sprintf("Hello from %s", os.Hostname())))
}

type healthCheckHandler struct{}

func (i healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
