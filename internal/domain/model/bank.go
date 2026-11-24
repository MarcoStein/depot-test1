package model

type Bank struct {
	BankAccounts []BankAccount
}

func (b *Bank) openBankAccount() error
func (b *Bank) closeBankAccount() error
func (b *Bank) rebookSecurity() error
