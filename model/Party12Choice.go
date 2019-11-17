package model

// Identification of a person, an organisation or a financial institution.
type Party12Choice struct {

	// Identification of a person or an organisation.
	Party *PartyIdentification43 `xml:"Pty"`

	// Identification of a financial institution.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (p *Party12Choice) AddParty() *PartyIdentification43 {
	p.Party = new(PartyIdentification43)
	return p.Party
}

func (p *Party12Choice) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agent
}
