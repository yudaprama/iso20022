package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account11 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification49Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account11) AddIdentification() *AccountIdentification1 {
	a.Identification = new(AccountIdentification1)
	return a.Identification
}

func (a *Account11) AddAccountServicer() *PartyIdentification49Choice {
	a.AccountServicer = new(PartyIdentification49Choice)
	return a.AccountServicer
}
