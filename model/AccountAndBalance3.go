package model

// Provides account and balance information.
type AccountAndBalance3 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails2 `xml:"Bal"`
}

func (a *AccountAndBalance3) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance3) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance3) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance3) AddBalance() *CorporateActionBalanceDetails2 {
	a.Balance = new(CorporateActionBalanceDetails2)
	return a.Balance
}
