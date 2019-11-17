package model

// Provides account and balance information.
type AccountAndBalance17 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails12 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance17) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance17) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance17) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance17) AddBalance() *CorporateActionBalanceDetails12 {
	a.Balance = new(CorporateActionBalanceDetails12)
	return a.Balance
}
