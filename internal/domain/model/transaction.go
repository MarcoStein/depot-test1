package model

import (
	"time"
	"uuid"
)

type Transaction struct {
	Id           uuid.UUID
	performed_at time.Time
	SecurityId   uuid.UUID // the involved security
}
