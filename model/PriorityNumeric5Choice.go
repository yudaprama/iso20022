package model

// Choice of format for the priority.
type PriorityNumeric5Choice struct {

	// Specifies the execution priority of the instruction with a number between 0001 and 9999.
	Numeric *Exact4NumericText `xml:"Nmrc"`

	// Specifies the execution priority of the instruction with a proprietary scheme.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *PriorityNumeric5Choice) SetNumeric(value string) {
	p.Numeric = (*Exact4NumericText)(&value)
}

func (p *PriorityNumeric5Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
