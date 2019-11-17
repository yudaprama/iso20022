package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account18 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification26 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification71Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account18) AddIdentification() *AccountIdentification26 {
	a.Identification = new(AccountIdentification26)
	return a.Identification
}

func (a *Account18) AddAccountServicer() *PartyIdentification71Choice {
	a.AccountServicer = new(PartyIdentification71Choice)
	return a.AccountServicer
}
