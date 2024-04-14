package routes

import (
	"github.com/kurjata/Project_Cadastro_de_Produtos_Golang/controllers"
	"net/http"
)


func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
