package model

// Nature or use of the account.
type Party13Choice struct {

	// Unique and unambiguous way to identify an organisation.
	OrganisationIdentification *OrganisationIdentification8 `xml:"OrgId"`

	// Unique and unambiguous identification of a financial institution.
	FinancialInstitutionIdentification *FinancialInstitutionIdentification9 `xml:"FIId"`
}

func (p *Party13Choice) AddOrganisationIdentification() *OrganisationIdentification8 {
	p.OrganisationIdentification = new(OrganisationIdentification8)
	return p.OrganisationIdentification
}

func (p *Party13Choice) AddFinancialInstitutionIdentification() *FinancialInstitutionIdentification9 {
	p.FinancialInstitutionIdentification = new(FinancialInstitutionIdentification9)
	return p.FinancialInstitutionIdentification
}
