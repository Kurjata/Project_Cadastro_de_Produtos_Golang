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

func Edit(w http.ResponseWriter, r *http.Request) {
    idProduto := r.URL.Query().Get("id")
    produto := models.EditarProduto(idProduto)
    temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        id := r.FormValue("id")
        nome := r.FormValue("nome")
        descricao := r.FormValue("descricao")
        preco := r.FormValue("preco")
        quantidade := r.FormValue("quantidade")

        idConvertido, err := strconv.Atoi(id)
        if err != nil {
            log.Println("Erro na conversão do ID para int:", err)
        }
        
        precoConvertido, err := strconv.ParseFloat(preco, 64)
        if err != nil {
            log.Println("Erro na conversão do preço para float:", err)
        }

        quantidadeConvertida, err := strconv.Atoi(quantidade)
        if err != nil {
            log.Println("Erro na conversão da quantidade para int:", err)
        }

        models.AtualizarProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConvertida)
    }

    http.Redirect(w, r, "/", 301)
    
}

func Delete(w http.ResponseWriter, r *http.Request) {
    idProduto := r.URL.Query().Get("id")
    models.DeletarProduto(idProduto)
    http.Redirect(w, r, "/", 301)
}