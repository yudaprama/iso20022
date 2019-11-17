package model

// Specifies rules governing an undertaking such as a guarantee or standby letter of credit.
type GovernanceRules2 struct {

	// Local identification to be used in IDREFs.
	Identification *ID `xml:"Id"`

	// Identification of the governance rules.
	RuleIdentification *GovernanceIdentification1Choice `xml:"RuleId"`

	// Law applicable to the undertaking.
	ApplicableLaw *Location1 `xml:"AplblLaw,omitempty"`

	// Place at or system under which any dispute related to the undertaking is to be resolved, such as court or arbitration. This is also known as 'forum'.
	Jurisdiction []*Location1 `xml:"Jursdctn,omitempty"`
}

func (g *GovernanceRules2) SetIdentification(value string) {
	g.Identification = (*ID)(&value)
}

func (g *GovernanceRules2) AddRuleIdentification() *GovernanceIdentification1Choice {
	g.RuleIdentification = new(GovernanceIdentification1Choice)
	return g.RuleIdentification
}

func (g *GovernanceRules2) AddApplicableLaw() *Location1 {
	g.ApplicableLaw = new(Location1)
	return g.ApplicableLaw
}

func (g *GovernanceRules2) AddJurisdiction() *Location1 {
	newValue := new(Location1)
	g.Jurisdiction = append(g.Jurisdiction, newValue)
	return newValue
}
