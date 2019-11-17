package model

// Legally constituted organization specified for this party.
type LegalOrganisation2 struct {

	// Unique and unambiguous identifier for an organisation that is allocated by an institution.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Specifies the short name of the organisation.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Date when the organisation was established.
	EstablishmentDate *ISODate `xml:"EstblishmtDt,omitempty"`

	// Date when the organisation was registered.
	RegistrationDate *ISODate `xml:"RegnDt,omitempty"`
}

func (l *LegalOrganisation2) SetIdentification(value string) {
	l.Identification = (*Max35Text)(&value)
}

func (l *LegalOrganisation2) SetName(value string) {
	l.Name = (*Max140Text)(&value)
}

func (l *LegalOrganisation2) SetEstablishmentDate(value string) {
	l.EstablishmentDate = (*ISODate)(&value)
}

func (l *LegalOrganisation2) SetRegistrationDate(value string) {
	l.RegistrationDate = (*ISODate)(&value)
}
