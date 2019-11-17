package model

// Identification of a person, an organisation or a financial institution.
type Party9Choice struct {

	// Identification of a person or an organisation.
	OrganisationIdentification *PartyIdentification42 `xml:"OrgId"`

	// Identification of a financial institution.
	FinancialInstitutionIdentification *BranchAndFinancialInstitutionIdentification5 `xml:"FIId"`
}

func (p *Party9Choice) AddOrganisationIdentification() *PartyIdentification42 {
	p.OrganisationIdentification = new(PartyIdentification42)
	return p.OrganisationIdentification
}

func (p *Party9Choice) AddFinancialInstitutionIdentification() *BranchAndFinancialInstitutionIdentification5 {
	p.FinancialInstitutionIdentification = new(BranchAndFinancialInstitutionIdentification5)
	return p.FinancialInstitutionIdentification
}
