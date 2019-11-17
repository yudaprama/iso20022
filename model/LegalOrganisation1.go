package model

// Legally constituted organization specified for this party.
type LegalOrganisation1 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution, eg, Dun & Bradstreet Identification.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Specifies the short name of the organisation.
	Name *Max140Text `xml:"Nm,omitempty"`
}

func (l *LegalOrganisation1) SetIdentification(value string) {
	l.Identification = (*Max35Text)(&value)
}

func (l *LegalOrganisation1) SetName(value string) {
	l.Name = (*Max140Text)(&value)
}
