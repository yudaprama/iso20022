package model

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification25 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification36Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance7 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification25) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification25) AddAccountOwner() *PartyIdentification36Choice {
	a.AccountOwner = new(PartyIdentification36Choice)
	return a.AccountOwner
}

func (a *AccountIdentification25) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification25) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance7 {
	newValue := new(CorporateActionEventAndBalance7)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}
