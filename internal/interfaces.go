package internal

import "time"

type IBank interface {
	OpenBankAccount(OpenAccountCommand) (BankAccount, error)
	CloseBankAccount()
	RebookSecurity()
}

type IBankAccount interface {
	OpenAccount() (Account, error)
	CloseAccount() error
	OpenDepot(OpenDepotCommand) (Depot, error)
	CloseDepot(CloseDepotCommand) error
	BankAccountValue(time.Time)
}

type IAccount interface {
	AccountValue(time.Time) error
}

type IDepot interface {
	BuySecurity() (Transaction, error)
	SellSecurity() (Transaction, error)
	DepotPosition(Isin string, PointInTime time.Time) (DepotPosition, error)
	// PositionValue(Isin string, PointInTime time.Time)
	TotalDepotValueCurrent()
	TotalDepotValue(time.Time)
	TotalDepotValueChange(time.Time)
	RebookSecurity()
}

type IDepotPosition interface {
	Amount() float64
	AmountAtTime(PointInTime time.Time)
}

type ISecurity interface {
}
