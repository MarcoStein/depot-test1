package model

import (
	"time"
	"uuid"
)

type Account struct {
	Id            uuid.UUID
	Name          string // Girokonto, etc
	AccountNumber string
}

func (a *Account) accountValue(time.Time) error
