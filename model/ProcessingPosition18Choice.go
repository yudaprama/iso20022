package model

// Choice of format for the processing position.
type ProcessingPosition18Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition4Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *ProcessingPosition18Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition4Code)(&value)
}

func (p *ProcessingPosition18Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
