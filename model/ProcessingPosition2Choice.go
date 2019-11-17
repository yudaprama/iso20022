package model

// Choice of format for the processing position.
type ProcessingPosition2Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition1Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *ProcessingPosition2Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition1Code)(&value)
}

func (p *ProcessingPosition2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
