package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router = httprouter.New()

	return router
}

func main() {
	r = httprouter.New()
	r.ListenAndServe(":9000")
}