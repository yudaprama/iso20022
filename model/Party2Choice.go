package model

// Choice between the identification of a person and the identification of a non-financial institution.
type Party2Choice struct {

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *OrganisationIdentification2 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification []*PersonIdentification3 `xml:"PrvtId"`
}

func (p *Party2Choice) AddOrganisationIdentification() *OrganisationIdentification2 {
	p.OrganisationIdentification = new(OrganisationIdentification2)
	return p.OrganisationIdentification
}

func (p *Party2Choice) AddPrivateIdentification() *PersonIdentification3 {
	newValue := new(PersonIdentification3)
	p.PrivateIdentification = append(p.PrivateIdentification, newValue)
	return newValue
}
