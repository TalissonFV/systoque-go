package models

import "fmt"

type Item struct {
	ID         int
	Nome       string
	Quantidade int
	Preco      float64
}

func (i *Item) Info() string {
	return fmt.Sprintf("ID: %d, Nome: %s, Quantidade: %d, Preço: %.2f", i.ID, i.Nome, i.Quantidade, i.Preco)
}
