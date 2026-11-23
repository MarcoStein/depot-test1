package internal

import (
	"time"
	"uuid"
)

type Bank struct {
	BankAccounts []BankAccount
}

type BankAccount struct {
	Owner    uuid.UUID
	Name     string
	Accounts []Account
	Depots   []Depot
}

type BankAccountOwner struct {
	Id    uuid.UUID
	Name  string
	Short string
}

type Account struct {
	Id            uuid.UUID
	AccountNumber string // 61883875
	Name          string // Girokonto, etc
}

type Depot struct {
	Id             uuid.UUID
	Transactions   []Transaction
	DepotPositions []DepotPosition
}

type Security struct {
	Id   uuid.UUID
	Isin string
	Gics GICS
}

type Transaction struct {
	Id           uuid.UUID
	performed_at time.Time
	SecurityId   uuid.UUID // the involved security
}

type DepotPosition struct {
}

// https://en.wikipedia.org/wiki/Global_Industry_Classification_Standard
type GICS struct {
}

type OpenBankAccount struct {
	Owner uuid.UUID
}

type OpenAccountCommand struct {
}

type OpenDepotCommand struct {
}

type CloseDepotCommand struct {
	DepotToBeClosed uuid.UUID
}

/* Depot structure. MUST BE refactored !
type Depot struct {
	gorm.Model
	Name string
	// Transactions Transactions
	traService       transaction.ITransactionService
	secService       securities.ISecurityService
	PortfolioService portfolio.IPortfolioService
}
*/

// ##################################
// --- STRING and PRINT functions ---
// func (d Depot) String() string { return d.Name }
// func (d Depot) Print()         { lslog.Debug("%s\n", d.String()) }
