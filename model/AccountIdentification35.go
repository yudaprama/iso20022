package model

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification35 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *RestrictedFINXMax35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification103Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat11Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance10 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification35) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*RestrictedFINXMax35Text)(&value)
}

func (a *AccountIdentification35) AddAccountOwner() *PartyIdentification103Choice {
	a.AccountOwner = new(PartyIdentification103Choice)
	return a.AccountOwner
}

func (a *AccountIdentification35) AddSafekeepingPlace() *SafekeepingPlaceFormat11Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat11Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification35) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance10 {
	newValue := new(CorporateActionEventAndBalance10)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}
