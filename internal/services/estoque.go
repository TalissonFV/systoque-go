package services

import (
	"fmt"
	"strconv"
	"systoque/internal/models"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  make([]models.Log, 0),
	}
}

func (e *Estoque) AdicionarItem(item models.Item) error {
	if item.Quantidade <= 0 {
		return fmt.Errorf("a quantidade do item deve ser maior que zero")
	}

	itemExistente, ok := e.items[strconv.Itoa(item.ID)]
	if ok {
		itemExistente.Quantidade += item.Quantidade
	}

	e.items[strconv.Itoa(item.ID)] = item
	log := models.Log{
		Timestamp:  time.Now(),
		Acao:       "Adicionado",
		ItemID:     item.ID,
		Quantidade: item.Quantidade,
		Razao:      "Item adicionado ao estoque",
	}
	e.logs = append(e.logs, log)
	return nil
}

func (e *Estoque) RemoverItem(nome string) {
	item, ok := e.items[nome]
	if ok {
		delete(e.items, nome)
		log := models.Log{
			Timestamp:  time.Now(),
			Acao:       "Removido",
			ItemID:     item.ID,
			Quantidade: item.Quantidade,
			Razao:      "Item removido do estoque",
		}
		e.logs = append(e.logs, log)
	}
}

func (e *Estoque) GetItem(nome string) (models.Item, bool) {
	item, ok := e.items[nome]
	return item, ok
}

func (e *Estoque) ListarItens() []models.Item {
	itens := []models.Item{}
	for _, item := range e.items {
		itens = append(itens, item)
	}
	return itens
}

func (e *Estoque) AdicionarQuantidade(nome string, quantidade int) {
	item, ok := e.items[nome]
	if ok {
		item.Quantidade += quantidade
		e.items[nome] = item
		log := models.Log{
			Timestamp:  time.Now(),
			Acao:       "Adicionado Quantidade",
			ItemID:     item.ID,
			Quantidade: quantidade,
			Razao:      fmt.Sprintf("%d unidades adicionadas ao item %s", quantidade, nome),
		}
		e.logs = append(e.logs, log)
	}
}

func (e *Estoque) RemoverQuantidade(nome string, quantidade int) {
	item, ok := e.items[nome]
	if ok {
		if item.Quantidade >= quantidade {
			item.Quantidade -= quantidade
			e.items[nome] = item
			log := models.Log{
				Timestamp:  time.Now(),
				Acao:       "Removido Quantidade",
				ItemID:     item.ID,
				Quantidade: quantidade,
				Razao:      fmt.Sprintf("%d unidades removidas do item %s", quantidade, nome),
			}
			e.logs = append(e.logs, log)
		}
	}
}

func (e *Estoque) ListarLogs() []models.Log {
	return e.logs
}
