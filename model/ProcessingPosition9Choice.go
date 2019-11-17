package model

// Choice of format for the processing position.
type ProcessingPosition9Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition5Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *ProcessingPosition9Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition5Code)(&value)
}

func (p *ProcessingPosition9Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
