package main

import (
	"log"
	"net/http"

	"github.com/waWanjohi/mux-auth/api/loggers"
	"github.com/waWanjohi/mux-auth/api/routers"
)

func main() {
	router := routers.MainRouter()
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	log.Fatal(http.ListenAndServe(":8080", loggers.RequestLogger(router)))
}
