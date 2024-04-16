package controllers

import (
    "html/template"
    "net/http"
    "github.com/kurjata/Project_Cadastro_de_Produtos_Golang/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
    todosProdutos := models.BuscaTodosProdutos()
    temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
    temp.ExecuteTemplate(w, "New", nil)
}
