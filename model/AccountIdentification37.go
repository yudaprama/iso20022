package model

// Provides account identification information.
type AccountIdentification37 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Balance to which the payment applies (less or equal to the total eligible balance).
	ConfirmedBalance *BalanceFormat7Choice `xml:"ConfdBal"`
}

func (a *AccountIdentification37) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification37) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification37) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification37) AddConfirmedBalance() *BalanceFormat7Choice {
	a.ConfirmedBalance = new(BalanceFormat7Choice)
	return a.ConfirmedBalance
}
