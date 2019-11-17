package model

// Party that provides services to investors relating to financial products.
type Intermediary4 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account2 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Max35Text `xml:"Role,omitempty"`
}

func (i *Intermediary4) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary4) AddAccount() *Account2 {
	i.Account = new(Account2)
	return i.Account
}

func (i *Intermediary4) SetRole(value string) {
	i.Role = (*Max35Text)(&value)
}
