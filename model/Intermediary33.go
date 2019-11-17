package model

// Identification of a party and its role.
type Intermediary33 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	Identification *PartyIdentification70Choice `xml:"Id"`

	// Identification of the party with a Legal Entity Identifier. This is a code allocated to a party as described in ISO 17442 "Financial Services - Legal Entity Identifier (LEI)".
	LegalEntityIdentifier *LEIIdentifier `xml:"LglNttyIdr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	Account *Account20 `xml:"Acct,omitempty"`
}

func (i *Intermediary33) AddIdentification() *PartyIdentification70Choice {
	i.Identification = new(PartyIdentification70Choice)
	return i.Identification
}

func (i *Intermediary33) SetLegalEntityIdentifier(value string) {
	i.LegalEntityIdentifier = (*LEIIdentifier)(&value)
}

func (i *Intermediary33) AddAccount() *Account20 {
	i.Account = new(Account20)
	return i.Account
}
