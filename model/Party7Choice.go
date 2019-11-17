package model

// Identification of a person, an organisation or a financial institution.
type Party7Choice struct {

	// Identification of a person or an organisation.
	Party *PartyIdentification32 `xml:"Pty"`

	// Identification of a financial institution.
	Agent *BranchAndFinancialInstitutionIdentification4 `xml:"Agt"`
}

func (p *Party7Choice) AddParty() *PartyIdentification32 {
	p.Party = new(PartyIdentification32)
	return p.Party
}

func (p *Party7Choice) AddAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification4)
	return p.Agent
}
