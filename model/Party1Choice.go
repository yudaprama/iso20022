package model

// Choice between the identification of a person and the identification of a non-financial institution.
type Party1Choice struct {

	// Unique and unambiguous way of identifying an organisation.
	OrganisationIdentification *NonFinancialInstitutionIdentification1 `xml:"OrgId"`

	// Unique and unambiguous identification of a person, eg, passport.
	PrivateIdentification []*PersonIdentification2 `xml:"PrvtId"`
}

func (p *Party1Choice) AddOrganisationIdentification() *NonFinancialInstitutionIdentification1 {
	p.OrganisationIdentification = new(NonFinancialInstitutionIdentification1)
	return p.OrganisationIdentification
}

func (p *Party1Choice) AddPrivateIdentification() *PersonIdentification2 {
	newValue := new(PersonIdentification2)
	p.PrivateIdentification = append(p.PrivateIdentification, newValue)
	return newValue
}
