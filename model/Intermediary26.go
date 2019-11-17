package model

// Party that provides services to investors relating to financial products.
type Intermediary26 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account14 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *Role4Choice `xml:"Role,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary26) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary26) AddAccount() *Account14 {
	i.Account = new(Account14)
	return i.Account
}

func (i *Intermediary26) AddRole() *Role4Choice {
	i.Role = new(Role4Choice)
	return i.Role
}

func (i *Intermediary26) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}
