package routes

import (
	"net/http"

	"github.com/igorferrati/servidor-go/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
