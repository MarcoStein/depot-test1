package internal

import (
	"time"
)

type IBank interface {
	OpenBankAccount(OpenAccountCommand) (IBankAccount, error)
	CloseBankAccount()
	RebookSecurity()
}

type IBankAccount interface {
	OpenAccount() (IAccount, error)
	CloseAccount() error
	OpenDepot(OpenDepotCommand) (IDepot, error)
	CloseDepot(CloseDepotCommand) error
	BankAccountValue(time.Time)
}

type IAccount interface {
	AccountValue(time.Time) error
}

type IDepot interface {
	BuySecurity() (ITransactionEvent, error)
	SellSecurity() (ITransactionEvent, error)
	DepotPosition(Isin string, PointInTime time.Time) (IDepotPosition, error)
	// PositionValue(Isin string, PointInTime time.Time)
	TotalDepotValueCurrent()
	TotalDepotValue(time.Time)
	TotalDepotValueChange(time.Time)
	RebookSecurity()
}

type ITransactionEvent interface {
}

type IDepotPosition interface {
	Amount() float64
	AmountAtTime(PointInTime time.Time)
}

type ISecurity interface {
}
