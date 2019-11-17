package model

// Choice of format for the processing position.
type ProcessingPosition10Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition3Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *ProcessingPosition10Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition3Code)(&value)
}

func (p *ProcessingPosition10Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
