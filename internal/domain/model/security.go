package model

import "uuid"

type Security struct {
	Id   uuid.UUID
	Isin string
	Gics GICS
}
