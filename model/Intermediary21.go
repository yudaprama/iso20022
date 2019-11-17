package model

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary21 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification49Choice `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role2Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account11 `xml:"Acct,omitempty"`
}

func (i *Intermediary21) AddIdentification() *PartyIdentification49Choice {
	i.Identification = new(PartyIdentification49Choice)
	return i.Identification
}

func (i *Intermediary21) AddRole() *Role2Choice {
	i.Role = new(Role2Choice)
	return i.Role
}

func (i *Intermediary21) AddAccount() *Account11 {
	i.Account = new(Account11)
	return i.Account
}
