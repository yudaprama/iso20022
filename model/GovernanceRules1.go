package model

// Rules governing an undertaking such as a guarantee or standby letter of credit.
type GovernanceRules1 struct {

	// Identification of the governance rules.
	RuleIdentification *GovernanceIdentification1Choice `xml:"RuleId"`

	// Law applicable to the undertaking.
	ApplicableLaw *Location1 `xml:"AplblLaw,omitempty"`

	// Place at or system under which any dispute related to the undertaking is to be resolved, such as court or arbitration. This is also known as 'forum'.
	Jurisdiction []*Location1 `xml:"Jursdctn,omitempty"`
}

func (g *GovernanceRules1) AddRuleIdentification() *GovernanceIdentification1Choice {
	g.RuleIdentification = new(GovernanceIdentification1Choice)
	return g.RuleIdentification
}

func (g *GovernanceRules1) AddApplicableLaw() *Location1 {
	g.ApplicableLaw = new(Location1)
	return g.ApplicableLaw
}

func (g *GovernanceRules1) AddJurisdiction() *Location1 {
	newValue := new(Location1)
	g.Jurisdiction = append(g.Jurisdiction, newValue)
	return newValue
}
