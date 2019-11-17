package model

// Provides account and balance information.
type AccountAndBalance36 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat5Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance36) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance36) AddConfirmedBalance() *BalanceFormat5Choice {
	a.ConfirmedBalance = new(BalanceFormat5Choice)
	return a.ConfirmedBalance
}
