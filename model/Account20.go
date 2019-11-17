package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account20 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Institution servicing an account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr"`
}

func (a *Account20) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account20) AddAccountServicer() *PartyIdentification70Choice {
	a.AccountServicer = new(PartyIdentification70Choice)
	return a.AccountServicer
}
