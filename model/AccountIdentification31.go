package model

// Provides account identification information.
type AccountIdentification31 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification31) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification31) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification31) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}
