package model

// Choice of format for the processing position.
type ProcessingPosition8Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition4Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *ProcessingPosition8Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition4Code)(&value)
}

func (p *ProcessingPosition8Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
