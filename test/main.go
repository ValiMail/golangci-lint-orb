package main

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.HandleFunc("/", RootHandler)
	http.Handle("/", r)
	_ = http.ListenAndServe("0.0.0.0:8080", r)
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	_, _ = io.WriteString(writer, "Hello World!")
	logrus.Info(request.RemoteAddr)
}
