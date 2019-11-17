package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account22 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`

	// Institution servicing the account and assigning the account identifier to the account owner.
	AccountServicer *PartyIdentification113 `xml:"AcctSvcr,omitempty"`
}

func (a *Account22) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *Account22) AddAccountServicer() *PartyIdentification113 {
	a.AccountServicer = new(PartyIdentification113)
	return a.AccountServicer
}
