package model

import "uuid"

type BankAccountOwner struct {
	Id    uuid.UUID
	Name  string
	Short string
}
