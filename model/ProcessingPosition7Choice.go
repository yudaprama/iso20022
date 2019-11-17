package model

// Choice of format for the processing position.
type ProcessingPosition7Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition3Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (p *ProcessingPosition7Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition3Code)(&value)
}

func (p *ProcessingPosition7Choice) AddProprietary() *GenericIdentification30 {
	p.Proprietary = new(GenericIdentification30)
	return p.Proprietary
}
