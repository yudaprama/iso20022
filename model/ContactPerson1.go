package model

// Contains information about the contact responsible for the transaction identified in the message.
type ContactPerson1 struct {

	// Provides additional information regarding the party, eg, the contact unit or person responsible for the transaction identified in the message.
	ContactPerson *ContactIdentification4 `xml:"CtctPrsn"`

	// Identification of the institution that the contact person represents.
	InstitutionIdentification *PartyIdentification2Choice `xml:"InstnId,omitempty"`
}

func (c *ContactPerson1) AddContactPerson() *ContactIdentification4 {
	c.ContactPerson = new(ContactIdentification4)
	return c.ContactPerson
}

func (c *ContactPerson1) AddInstitutionIdentification() *PartyIdentification2Choice {
	c.InstitutionIdentification = new(PartyIdentification2Choice)
	return c.InstitutionIdentification
}
