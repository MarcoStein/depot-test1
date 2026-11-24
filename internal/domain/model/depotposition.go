package model

import "time"

type DepotPosition struct {
}

func (dp *DepotPosition) positionValue(Isin string, PointInTime time.Time) error
func (dp *DepotPosition) amount() float64
func (dp *DepotPosition) amountAtTime(PointInTime time.Time)
