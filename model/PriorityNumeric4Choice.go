package model

// Choice of format for the priority.
type PriorityNumeric4Choice struct {

	// Specifies the execution priority of the instruction with a number between 0001 and 9999.
	Numeric *Exact4NumericText `xml:"Nmrc"`

	// Specifies the execution priority of the instruction with a proprietary scheme.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *PriorityNumeric4Choice) SetNumeric(value string) {
	p.Numeric = (*Exact4NumericText)(&value)
}

func (p *PriorityNumeric4Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
