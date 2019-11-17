package model

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary23 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account13 `xml:"Acct,omitempty"`
}

func (i *Intermediary23) AddIdentification() *PartyIdentification49Choice {
	i.Identification = new(PartyIdentification49Choice)
	return i.Identification
}

func (i *Intermediary23) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary23) AddAccount() *Account13 {
	i.Account = new(Account13)
	return i.Account
}
