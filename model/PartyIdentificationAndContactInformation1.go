package model

// Identification of a person, or a non-financial institution.
type PartyIdentificationAndContactInformation1 struct {

	// Identification of the party.
	PartyIdentification *PartyIdentification8 `xml:"PtyId"`

	// Information needed to contact a physical person related to the party, such as name, phone number, email address.
	ContactInformation *ContactIdentification1 `xml:"CtctInf,omitempty"`
}

func (p *PartyIdentificationAndContactInformation1) AddPartyIdentification() *PartyIdentification8 {
	p.PartyIdentification = new(PartyIdentification8)
	return p.PartyIdentification
}

func (p *PartyIdentificationAndContactInformation1) AddContactInformation() *ContactIdentification1 {
	p.ContactInformation = new(ContactIdentification1)
	return p.ContactInformation
}
