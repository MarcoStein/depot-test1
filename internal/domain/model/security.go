package model

import "uuid"

type Security struct {
	Id   uuid.UUID
	Isin string // Should the ISIN be the ultimate identifier
	Gics GICS
}
