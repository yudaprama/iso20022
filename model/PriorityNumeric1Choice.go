package model

// Choice of format for the priority.
type PriorityNumeric1Choice struct {

	// Specifies the execution priority of the instruction with a number between 0001 and 9999.
	Numeric *Exact4NumericText `xml:"Nmrc"`

	// Specifies the execution priority of the instruction with a proprietary scheme.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PriorityNumeric1Choice) SetNumeric(value string) {
	p.Numeric = (*Exact4NumericText)(&value)
}

func (p *PriorityNumeric1Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
