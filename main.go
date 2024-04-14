package main

import (
	"net/http"
	"github.com/kurjata/Project_Cadastro_de_Produtos_Golang/routes"
)


func main() {
	routes.LoadRoutes()	
    http.ListenAndServe(":8000", nil)
}


