package model

// Party that provides services to investors relating to financial products.
type Intermediary27 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`
}

func (i *Intermediary27) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary27) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary27) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}
