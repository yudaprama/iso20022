package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account7 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account7) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account7) AddAccountServicer() *PartyIdentification2Choice {
	a.AccountServicer = new(PartyIdentification2Choice)
	return a.AccountServicer
}
