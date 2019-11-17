package model

// Account information and detailed account holdings information report for corporate action events.
type AccountIdentification33 struct {

	// Account where financial instruments are maintained.
	SafekeepingAccount *Max35Text `xml:"SfkpgAcct"`

	// Party that legally owns the account.
	AccountOwner *PartyIdentification92Choice `xml:"AcctOwnr,omitempty"`

	// Location where the financial instruments are/will be safekept.
	SafekeepingPlace *SafekeepingPlaceFormat8Choice `xml:"SfkpgPlc,omitempty"`

	// Detailed account holdings information report for a corporate action event.
	CorporateActionEventAndBalance []*CorporateActionEventAndBalance9 `xml:"CorpActnEvtAndBal,omitempty"`
}

func (a *AccountIdentification33) SetSafekeepingAccount(value string) {
	a.SafekeepingAccount = (*Max35Text)(&value)
}

func (a *AccountIdentification33) AddAccountOwner() *PartyIdentification92Choice {
	a.AccountOwner = new(PartyIdentification92Choice)
	return a.AccountOwner
}

func (a *AccountIdentification33) AddSafekeepingPlace() *SafekeepingPlaceFormat8Choice {
	a.SafekeepingPlace = new(SafekeepingPlaceFormat8Choice)
	return a.SafekeepingPlace
}

func (a *AccountIdentification33) AddCorporateActionEventAndBalance() *CorporateActionEventAndBalance9 {
	newValue := new(CorporateActionEventAndBalance9)
	a.CorporateActionEventAndBalance = append(a.CorporateActionEventAndBalance, newValue)
	return newValue
}
