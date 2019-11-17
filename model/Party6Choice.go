package model

// Nature or use of the account.
type Party6Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification4 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification *PersonIdentification5 `xml:"PrvtId"`
}

func (p *Party6Choice) AddOrganisationIdentification() *OrganisationIdentification4 {
	p.OrganisationIdentification = new(OrganisationIdentification4)
	return p.OrganisationIdentification
}

func (p *Party6Choice) AddPrivateIdentification() *PersonIdentification5 {
	p.PrivateIdentification = new(PersonIdentification5)
	return p.PrivateIdentification
}
