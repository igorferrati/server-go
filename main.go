package main

import (
	"net/http"

	"github.com/igorferrati/servidor-go/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.Routes()
	http.ListenAndServe(":8001", nil)
}
