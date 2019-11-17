package model

// Nature or use of the account.
type Party10Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification7 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification *PersonIdentification5 `xml:"PrvtId"`
}

func (p *Party10Choice) AddOrganisationIdentification() *OrganisationIdentification7 {
	p.OrganisationIdentification = new(OrganisationIdentification7)
	return p.OrganisationIdentification
}

func (p *Party10Choice) AddPrivateIdentification() *PersonIdentification5 {
	p.PrivateIdentification = new(PersonIdentification5)
	return p.PrivateIdentification
}
