package model

// Nature or use of the account.
type Party11Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification *PersonIdentification5 `xml:"PrvtId"`
}

func (p *Party11Choice) AddOrganisationIdentification() *OrganisationIdentification8 {
	p.OrganisationIdentification = new(OrganisationIdentification8)
	return p.OrganisationIdentification
}

func (p *Party11Choice) AddPrivateIdentification() *PersonIdentification5 {
	p.PrivateIdentification = new(PersonIdentification5)
	return p.PrivateIdentification
}
