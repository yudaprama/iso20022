package model

// Choice of format for the processing position.
type ProcessingPosition5Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition4Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *ProcessingPosition5Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition4Code)(&value)
}

func (p *ProcessingPosition5Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
