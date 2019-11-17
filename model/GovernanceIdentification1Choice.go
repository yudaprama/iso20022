package model

// Choice of format for the governance rules.
type GovernanceIdentification1Choice struct {

	// Governance identification information.
	//
	Code *GovernanceIdentification1Code `xml:"Cd"`

	// Governance identification information expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (g *GovernanceIdentification1Choice) SetCode(value string) {
	g.Code = (*GovernanceIdentification1Code)(&value)
}

func (g *GovernanceIdentification1Choice) AddProprietary() *GenericIdentification1 {
	g.Proprietary = new(GenericIdentification1)
	return g.Proprietary
}
