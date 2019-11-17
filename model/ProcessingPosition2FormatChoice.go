package model

// Choice of formats to  express the processing position.
type ProcessingPosition2FormatChoice struct {

	// Standard code to specify the processing position.
	Code *ProcessingPosition2Code `xml:"Cd"`

	// Proprietary code to  express the processing position.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessingPosition2FormatChoice) SetCode(value string) {
	p.Code = (*ProcessingPosition2Code)(&value)
}

func (p *ProcessingPosition2FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
