package model

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification6 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat2Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance1 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification6) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification6) AddAccountOwner() *PartyIdentification13Choice {
	a.AccountOwner = new(PartyIdentification13Choice)
	return a.AccountOwner
}

func (a *AccountIdentification6) AddSafekeepingPlace() *SafekeepingPlaceFormat2Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat2Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification6) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance1 {
	newValue := new(CorporateActionEventAndBalance1)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}
