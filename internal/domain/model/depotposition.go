package model

import (
	"time"
	"uuid"
)

type DepotPosition struct {
	Amount     float64
	SecurityId uuid.UUID // references the UUID not the ISIN
}

type DepotPositions []DepotPosition

func (dp *DepotPosition) sellPositionCompletely() error
func (dp *DepotPosition) sellPositionPartly() error // amount float64
func (dp *DepotPosition) increasePosition() error   // min amount float64 as parameter

func (dp *DepotPosition) positionValue(PointInTime time.Time) error
func (dp *DepotPosition) amount() float64
func (dp *DepotPosition) amountAtTime(PointInTime time.Time)
