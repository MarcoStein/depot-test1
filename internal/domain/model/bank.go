package model

type Bank struct {
	bankAccounts BankAccounts
	isin         string
}

func (b *Bank) openBankAccount() error
func (b *Bank) closeBankAccount() error
func (b *Bank) rebookSecurity() error
