package model

// Ancillary identification information about the party.
type PartyAdditionalIdentification2Choice struct {

	// Date of birth of an individual person.
	BirthDate *ISODate `xml:"BirthDt"`

	// Official identification of an organisation in a specific register.
	RegistrationIdentification *OrganisationIdentification5 `xml:"RegnId"`
}

func (p *PartyAdditionalIdentification2Choice) SetBirthDate(value string) {
	p.BirthDate = (*ISODate)(&value)
}

func (p *PartyAdditionalIdentification2Choice) AddRegistrationIdentification() *OrganisationIdentification5 {
	p.RegistrationIdentification = new(OrganisationIdentification5)
	return p.RegistrationIdentification
}
