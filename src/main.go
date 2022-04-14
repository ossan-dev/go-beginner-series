package main

import (
	"fmt"
	"net/http"

	"todoapi/router"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starting up HTTP server...")

	r := mux.NewRouter()

	router.SetupRoutes(r)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic(err)
	}
}
