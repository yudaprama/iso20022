package model

// Nature or use of the account.
type Party8Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification6 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification *PersonIdentification5 `xml:"PrvtId"`
}

func (p *Party8Choice) AddOrganisationIdentification() *OrganisationIdentification6 {
	p.OrganisationIdentification = new(OrganisationIdentification6)
	return p.OrganisationIdentification
}

func (p *Party8Choice) AddPrivateIdentification() *PersonIdentification5 {
	p.PrivateIdentification = new(PersonIdentification5)
	return p.PrivateIdentification
}
