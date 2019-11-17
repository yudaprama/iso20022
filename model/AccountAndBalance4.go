package model

// Provides account and balance information.
type AccountAndBalance4 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat1Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance4) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance4) AddConfirmedBalance() *BalanceFormat1Choice {
	a.ConfirmedBalance = new(BalanceFormat1Choice)
	return a.ConfirmedBalance
}
