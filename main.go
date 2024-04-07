package main

import (
	"net/http"
	"text/template"
)

type Produto struct{

	Nome string
	Descricao string
	Preco float64
	Quantidade int

}
	

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Camiseta preta", Preco: 39.90, Quantidade: 10},
		{"Tênis", "Air Jordan", 999, 5},
		{"Bermuda", "Bermuda de praia", 59.90, 15},
		{"Relógio", "Relógio de pulso", 299.90, 3},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}