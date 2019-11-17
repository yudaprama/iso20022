package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type Account21 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification4 `xml:"Id"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification104Choice `xml:"AcctSvcr,omitempty"`
}

func (a *Account21) AddIdentification() *AccountIdentification4 {
	a.Identification = new(AccountIdentification4)
	return a.Identification
}

func (a *Account21) AddAccountServicer() *PartyIdentification104Choice {
	a.AccountServicer = new(PartyIdentification104Choice)
	return a.AccountServicer
}
