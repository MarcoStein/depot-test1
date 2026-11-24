package model

import (
	"time"
	"uuid"
)

type BackAccounts []BankAccount

type BankAccount struct {
	Id       uuid.UUID
	OwnerId  uuid.UUID // ref BankAccountOwner
	Name     string
	Accounts []Account
	Depots   []Depot
}

func (ba *BankAccount) openAccount() error
func (ba *BankAccount) closeAccount() error
func (ba *BankAccount) openDepot() error
func (ba *BankAccount) closeDepot() error
func (ba *BankAccount) bankAccountValue(time.Time) error
