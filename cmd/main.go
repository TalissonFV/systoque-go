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
		if err := estoque.AdicionarItem(item, "admin"); err != nil {
			fmt.Println(err)
			continue
		}
	}

	for _, item := range estoque.ListarItens() {
		fmt.Println(item.Info())
	}

	for _, log := range estoque.ListarLogs() {
		fmt.Printf("[%s] %s, itemID: %d, por: %s, motivo: %s, quantidade: %d\n", log.Timestamp.Format("2006-01-02 15:04:05"), log.Acao, log.ItemID, log.User, log.Razao, log.Quantidade)
	}

}
