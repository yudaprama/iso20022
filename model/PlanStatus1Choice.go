package model

// Choice of formats for the status of a plan.
type PlanStatus1Choice struct {

	// Status of the plan expressed as a code.
	Code *PlanStatus1Code `xml:"Cd"`

	// Status of the plan expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (p *PlanStatus1Choice) SetCode(value string) {
	p.Code = (*PlanStatus1Code)(&value)
}

func (p *PlanStatus1Choice) AddProprietary() *GenericIdentification29 {
	p.Proprietary = new(GenericIdentification29)
	return p.Proprietary
}
