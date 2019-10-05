package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/averagemarcus/chaos-gaming/internal/routes"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8000"
	}

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/api/{resource}/", routes.FetchResources).Methods(http.MethodGet)
	r.HandleFunc("/api/{resource}/{uid}/", routes.KillResource).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
