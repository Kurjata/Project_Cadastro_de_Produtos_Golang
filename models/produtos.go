package models

import (
	"github.com/kurjata/Project_Cadastro_de_Produtos_Golang/db"


)

type Produto struct {
    Id         int
    Nome       string
    Descricao  string
    Preco      float64
    Quantidade int
}

func BuscaTodosProdutos() []Produto {
    db := db.ConectDataBase() 
    selectTodosProdutos, err := db.Query("SELECT * FROM produtos ORDER BY id ASC")
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

        p.Id = id
        p.Nome = nome
        p.Descricao = descricao
        p.Preco = preco
        p.Quantidade = quantidade

        produtos = append(produtos, p)
    }

    defer db.Close()
    return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
    db := db.ConectDataBase()
    insertProduto, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1, $2, $3, $4)")
    if err != nil {
        panic(err.Error())
    }

    insertProduto.Exec(nome, descricao, preco, quantidade)
    defer db.Close()
}

func EditarProduto(id string) Produto {
    db := db.ConectDataBase()
    produtoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
    if err != nil {
        panic(err.Error())
    }

    produtoAtualizado := Produto{}

    for produtoBanco.Next() {
        var id, quantidade int
        var nome, descricao string
        var preco float64

        err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
        if err != nil {
            panic(err.Error())
        }

        produtoAtualizado.Id = id
        produtoAtualizado.Nome = nome
        produtoAtualizado.Descricao = descricao
        produtoAtualizado.Preco = preco
        produtoAtualizado.Quantidade = quantidade
    }
    defer db.Close()
    return produtoAtualizado
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
    db := db.ConectDataBase()
    atualizarProduto, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
    if err != nil {
        panic(err.Error())
    }

    atualizarProduto.Exec(nome, descricao, preco, quantidade, id)
    defer db.Close()
}

func DeletarProduto(id string) {
    db := db.ConectDataBase()
    deletarProduto, err := db.Prepare("DELETE FROM produtos WHERE id=$1")
    if err != nil {
        panic(err.Error())
    }

    deletarProduto.Exec(id)
    defer db.Close()
}