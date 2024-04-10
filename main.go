package main

import (
	"database/sql"
	"fmt"
	"os"
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
)

func conectDataBase() *sql.DB {

	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	sslmode := os.Getenv("DB_SSLMODE")

	conect := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s", user, dbname, password, host, sslmode)
	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct{
	Id int
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
	db := conectDataBase()
	selectTodosProdutos, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectTodosProdutos.Next() {

		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
	
}