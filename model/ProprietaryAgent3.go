package model

// Identifies a proprietary party.
type ProprietaryAgent3 struct {

	// Specifies the type of proprietary agent.
	Type *Max35Text `xml:"Tp"`

	// Organisation established primarily to provide financial services.
	Agent *BranchAndFinancialInstitutionIdentification5 `xml:"Agt"`
}

func (p *ProprietaryAgent3) SetType(value string) {
	p.Type = (*Max35Text)(&value)
}

func (p *ProprietaryAgent3) AddAgent() *BranchAndFinancialInstitutionIdentification5 {
	p.Agent = new(BranchAndFinancialInstitutionIdentification5)
	return p.Agent
}
