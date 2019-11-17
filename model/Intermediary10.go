package model

// Party that provides services to investors relating to financial products.
type Intermediary10 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`

	// Contact person and contact information.
	ContactPerson *ContactIdentification2 `xml:"CtctPrsn,omitempty"`
}

func (i *Intermediary10) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary10) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary10) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary10) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}

func (i *Intermediary10) AddContactPerson() *ContactIdentification2 {
	i.ContactPerson = new(ContactIdentification2)
	return i.ContactPerson
}
