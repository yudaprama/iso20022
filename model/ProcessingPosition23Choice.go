package model

// Choice of format for the processing position.
type ProcessingPosition23Choice struct {

	// Processing position expressed as an ISO 20022 code.
	Code *ProcessingPosition5Code `xml:"Cd"`

	// Processing position expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (p *ProcessingPosition23Choice) SetCode(value string) {
	p.Code = (*ProcessingPosition5Code)(&value)
}

func (p *ProcessingPosition23Choice) AddProprietary() *GenericIdentification47 {
	p.Proprietary = new(GenericIdentification47)
	return p.Proprietary
}
