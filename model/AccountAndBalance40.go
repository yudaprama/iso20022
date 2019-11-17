package model

// Provides account and balance information.
type AccountAndBalance40 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`
}

func (a *AccountAndBalance40) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance40) AddConfirmedBalance() *BalanceFormat7Choice {
	a.ConfirmedBalance = new(BalanceFormat7Choice)
	return a.ConfirmedBalance
}
