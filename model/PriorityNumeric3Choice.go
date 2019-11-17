package model

// Choice of format for the priority.
type PriorityNumeric3Choice struct {

	// Specifies the execution priority of the instruction with a number between 0001 and 9999.
	Numeric *Exact4NumericText `xml:"Nmrc"`

	// Specifies the execution priority of the instruction with a proprietary scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (p *PriorityNumeric3Choice) SetNumeric(value string) {
	p.Numeric = (*Exact4NumericText)(&value)
}

func (p *PriorityNumeric3Choice) AddProprietary() *GenericIdentification38 {
	p.Proprietary = new(GenericIdentification38)
	return p.Proprietary
}
