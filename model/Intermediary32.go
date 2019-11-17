package model

// Party that provides services to investors relating to financial products (Investment Funds).
type Intermediary32 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification100 `xml:"Id"`

	// Function performed by the intermediary (investment funds).
	Role *Role6Choice `xml:"Role,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account18 `xml:"Acct,omitempty"`
}

func (i *Intermediary32) AddIdentification() *PartyIdentification100 {
	i.Identification = new(PartyIdentification100)
	return i.Identification
}

func (i *Intermediary32) AddRole() *Role6Choice {
	i.Role = new(Role6Choice)
	return i.Role
}

func (i *Intermediary32) AddAccount() *Account18 {
	i.Account = new(Account18)
	return i.Account
}
