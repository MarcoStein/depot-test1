package model

import (
	"time"
	"uuid"
)

type BankAccount struct {
	Id       uuid.UUID
	OwnerId  uuid.UUID // ref BankAccountOwner
	Name     string
	Accounts []Account
	Depots   []Depot
}

type BankAccounts []BankAccount

func (ba *BankAccount) openAccount() error
func (ba *BankAccount) closeAccount() error
func (ba *BankAccount) openDepot() error
func (ba *BankAccount) closeDepot() error
func (ba *BankAccount) bankAccountValue(time.Time) error
