package model

// Provides account and balance information.
type AccountAndBalance33 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails29 `xml:"Bal,omitempty"`
}

func (a *AccountAndBalance33) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountAndBalance33) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance33) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance33) AddBalance() *CorporateActionBalanceDetails29 {
	a.Balance = new(CorporateActionBalanceDetails29)
	return a.Balance
}
