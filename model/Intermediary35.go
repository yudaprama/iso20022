package model

// Party that provides services to investors relating to financial products.
type Intermediary35 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`
}

func (i *Intermediary35) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary35) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}

func (i *Intermediary35) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}
