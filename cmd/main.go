package main

import (
	"fmt"
	"systoque/internal/models"
	"systoque/internal/services"
)

func main() {
	itens := []models.Item{
		{
			ID:         1,
			Nome:       "Item 1",
			Quantidade: 10,
			Preco:      10.99,
		},
		{
			ID:         2,
			Nome:       "Item 2",
			Quantidade: 5,
			Preco:      5.99,
		},
	}

	estoque := services.NewEstoque()

	for _, item := range itens {
		if err := estoque.AdicionarItem(item); err != nil {
			fmt.Println(err)
			continue
		}
	}


	for _,item := range estoque.ListarItens(){
		fmt.Println(item.Info())
	}

}
