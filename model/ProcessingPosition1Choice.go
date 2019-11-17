package model

// Choice of format for the processing position.
type ProcessingPosition1Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition3Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *ProcessingPosition1Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition3Code)(&value)
}

func (p *ProcessingPosition1Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
