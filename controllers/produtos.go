package controllers

import (
	"net/http"
	"text/template"

	"github.com/igorferrati/servidor-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	TodosProdutos := models.BuscaProdutos()
	temp.ExecuteTemplate(w, "Index", TodosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
