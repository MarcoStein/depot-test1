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
func (d *Depot) sellSecurity() error
func (d *Depot) depotPosition(Isin string, PointInTime time.Time) (DepotPosition, error)

func (d *Depot) totalDepotValueCurrent() error
func (d *Depot) totalDepotValue(time.Time) error
func (d *Depot) totalDepotValueChange(time.Time) error
func (d *Depot) rebookSecurity() error
