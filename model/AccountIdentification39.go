package model

// Provides account identification information.
type AccountIdentification39 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification95Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat26Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification39) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification39) AddAccountOwner() *PartyIdentification95Choice {
	a.AccountOwner = new(PartyIdentification95Choice)
	return a.AccountOwner
}

func (a *AccountIdentification39) AddSafekeepingPlace() *SafekeepingPlaceFormat26Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat26Choice)
	return a.SafekeepingPlace
}
