package model

import (
	"time"
	"uuid"
)

type Depot struct {
	Id             uuid.UUID
	Name           string
	DepotNumber    string
	Transactions   Transactions
	DepotPositions DepotPositions
}

func (d *Depot) buySecurity() error
func (d *Depot) sellSecurity() error // could be moved to the DepotPosition
// func (d *Depot) depotPosition(PointInTime time.Time) (DepotPosition, error) // <-- that Is-No-Good

func (d *Depot) totalDepotValueCurrent() error // convenience function for totalDepotValue(time.Now())
func (d *Depot) totalDepotValue(time.Time) error
func (d *Depot) totalDepotValueChange(time.Time) error
func (d *Depot) rebookSecurity() error
