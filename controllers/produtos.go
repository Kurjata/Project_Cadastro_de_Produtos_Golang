package controllers

import (
    "html/template"
    "net/http"
    "github.com/kurjata/Project_Cadastro_de_Produtos_Golang/models"
    "log"
    "strconv"

)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
    todosProdutos := models.BuscaTodosProdutos()
    temp.ExecuteTemplate(w, "Index", todosProdutos)
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
            log.Println("Erro na conversão do preço:", err)
        }

        quantidadeConvertida, err := strconv.Atoi(quantidade)
        if err != nil {
            log.Println("Erro na conversão da quantidade:", err)
        }

        models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
    }
    http.Redirect(w, r, "/", 301)
}