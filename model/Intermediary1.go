package model

// Organised structure that is set up for a particular purpose, eg, a business, government body, department, charity, or financial institution.
type Intermediary1 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *PartyIdentification1Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account1 `xml:"Acct,omitempty"`

	// Set of  functions performed by an intermediary in a given situation.
	Role *Max35Text `xml:"Role,omitempty"`
}

func (i *Intermediary1) AddIdentification() *PartyIdentification1Choice {
	i.Identification = new(PartyIdentification1Choice)
	return i.Identification
}

func (i *Intermediary1) AddAccount() *Account1 {
	i.Account = new(Account1)
	return i.Account
}

func (i *Intermediary1) SetRole(value string) {
	i.Role = (*Max35Text)(&value)
}
