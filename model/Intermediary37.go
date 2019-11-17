package model

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary37 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification111 `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role7Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account21 `xml:"Acct,omitempty"`
}

func (i *Intermediary37) AddIdentification() *PartyIdentification111 {
	i.Identification = new(PartyIdentification111)
	return i.Identification
}

func (i *Intermediary37) AddRole() *Role7Choice {
	i.Role = new(Role7Choice)
	return i.Role
}

func (i *Intermediary37) AddAccount() *Account21 {
	i.Account = new(Account21)
	return i.Account
}
