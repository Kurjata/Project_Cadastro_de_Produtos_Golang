package main

import (
    "html/template"
    "net/http"
    "github.com/kurjata/Project_Cadastro_de_Produtos_Golang/models" // Ajuste na importação
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
    http.HandleFunc("/", index)
    http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
    todosProdutos := models.BuscaTodosProdutos()
    temp.ExecuteTemplate(w, "Index", todosProdutos)
}
