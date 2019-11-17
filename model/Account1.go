package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id,omitempty"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr"`
}

func (a *Account1) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account1) AddAccountServicer() *PartyIdentification1Choice {
	a.AccountServicer = new(PartyIdentification1Choice)
	return a.AccountServicer
}
