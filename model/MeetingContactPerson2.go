package model

// Contact person at the party organising the meeting, at the issuer or at an intermediary.
type MeetingContactPerson2 struct {

	// Identifies the contact person by its name, given name and address.
	ContactPerson *ContactIdentification1 `xml:"CtctPrsn,omitempty"`

	// Identifies the organisation which is represented by the person or for which the person works.
	EmployingParty *PartyIdentification40Choice `xml:"EmplngPty,omitempty"`

	// Identification of the financial market, as stipulated in the norm ISO 10383 'Codes for exchanges and market identifications'.
	PlaceOfListing *MICIdentifier `xml:"PlcOfListg,omitempty"`
}

func (m *MeetingContactPerson2) AddContactPerson() *ContactIdentification1 {
	m.ContactPerson = new(ContactIdentification1)
	return m.ContactPerson
}

func (m *MeetingContactPerson2) AddEmployingParty() *PartyIdentification40Choice {
	m.EmployingParty = new(PartyIdentification40Choice)
	return m.EmployingParty
}

func (m *MeetingContactPerson2) SetPlaceOfListing(value string) {
	m.PlaceOfListing = (*MICIdentifier)(&value)
}
