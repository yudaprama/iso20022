package model

// Identification of a person, an organisation or a financial institution.
type Party28Choice struct {

	// Identification of a person or an organisation.
	Party *PartyIdentification77 `xml:"Pty"`

	// Identification of a financial institution.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (p *Party28Choice) AddParty() *PartyIdentification77 {
	p.Party = new(PartyIdentification77)
	return p.Party
}

func (p *Party28Choice) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agent
}
