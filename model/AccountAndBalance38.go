package model

// Provides account and balance information.
type AccountAndBalance38 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Provides information about balance related to a corporate action.
	Balance *CorporateActionBalanceDetails35 `xml:"Bal"`
}

func (a *AccountAndBalance38) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountAndBalance38) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountAndBalance38) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountAndBalance38) AddBalance() *CorporateActionBalanceDetails35 {
	a.Balance = new(CorporateActionBalanceDetails35)
	return a.Balance
}
