package model

// Set of elements to identify a proprietary party.
type ProprietaryAgent1 struct {

	// Identifies the type of proprietary agent reported.
	Type *Max35Text `xml:"Tp"`

	// Proprietary agent.
	Agent *BranchAndFinancialInstitutionIdentification3 `xml:"Agt"`
}

func (p *ProprietaryAgent1) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent1) AddAgent() *BranchAndFinancialInstitutionIdentification3 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification3)
	return p.Agent
}
