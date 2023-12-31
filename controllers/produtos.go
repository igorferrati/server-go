package controllers

import (
	"log"
	"net/http"
	"strconv"
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

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preço, erro: ", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter a quantidade, erro: ", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id") //pegar na url

	models.DeleteProduct(idProduto)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")

	produto := models.EditProduct(idProduto)

	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro ao converter id, erro: ", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade, erro: ", err)
		}

		precoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter o preço, erro: ", err)
		}

		models.AtualizarProduto(idInt, nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}
