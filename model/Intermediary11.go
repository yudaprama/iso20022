package model

// Party that provides services to investors relating to financial products.
type Intermediary11 struct {

	// Unique and unambiguous identifier of the intermediary.
	Identification *PartyIdentification2Choice `xml:"Id"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account7 `xml:"Acct,omitempty"`

	// Function performed by the intermediary.
	Role *InvestmentFundRole2Code `xml:"Role,omitempty"`

	// Function performed by the intermediary.
	ExtendedRole *Extended350Code `xml:"XtndedRole,omitempty"`
}

func (i *Intermediary11) AddIdentification() *PartyIdentification2Choice {
	i.Identification = new(PartyIdentification2Choice)
	return i.Identification
}

func (i *Intermediary11) AddAccount() *Account7 {
	i.Account = new(Account7)
	return i.Account
}

func (i *Intermediary11) SetRole(value string) {
	i.Role = (*InvestmentFundRole2Code)(&value)
}

func (i *Intermediary11) SetExtendedRole(value string) {
	i.ExtendedRole = (*Extended350Code)(&value)
}
