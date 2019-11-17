package model

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary2 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification10Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account9 `xml:"Acct,omitempty"`
}

func (i *Intermediary2) AddIdentification() *PartyIdentification10Choice {
	i.Identification = new(PartyIdentification10Choice)
	return i.Identification
}

func (i *Intermediary2) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary2) AddAccount() *Account9 {
	i.Account = new(Account9)
	return i.Account
}
