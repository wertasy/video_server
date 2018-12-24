package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", indexHandler)
	router.POST("/", loginHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)

	router.ServeFiles("/assets/*filepath", http.Dir("./templates"))
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}
