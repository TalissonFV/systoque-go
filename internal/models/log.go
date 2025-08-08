package models

import "time"

type Log struct {
	Timestamp  time.Time
	Acao       string
	User       string
	ItemID     int
	Quantidade int
	Razao      string
}
