package model

// Choice of formats for the status of a plan.
type PlanStatus2Choice struct {

	// Status of the plan expressed as a code.
	Code *PlanStatus1Code `xml:"Cd"`

	// Status of the plan expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PlanStatus2Choice) SetCode(value string) {
	p.Code = (*PlanStatus1Code)(&value)
}

func (p *PlanStatus2Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
