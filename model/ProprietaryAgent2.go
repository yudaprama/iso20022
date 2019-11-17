package model

// Set of elements used to identify a proprietary party.
type ProprietaryAgent2 struct {

	// Specifies the type of proprietary agent.
	Type *Max35Text `xml:"Tp"`

	// Organisation established primarily to provide financial services.
	Agent *BranchAndFinancialInstitutionIdentification4 `xml:"Agt"`
}

func (p *ProprietaryAgent2) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent2) AddAgent() *BranchAndFinancialInstitutionIdentification4 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification4)
	return p.Agent
}
