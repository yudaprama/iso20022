package model

// Provides account identification information.
type AccountIdentification42 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat27Choice `xml:"SfkpgPlc,omitempty"`
}

func (a *AccountIdentification42) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification42) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification42) AddSafekeepingPlace() *SafekeepingPlaceFormat27Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat27Choice)
	return a.SafekeepingPlace
}
