package model

// Choice of formats to  express the processing status of an advice, a request or a movement.
type ProcessedStatus3FormatChoice struct {

	// Standard code to specify the processing status of an advice, a request or a movement.
	Code *ProcessedStatus3Code `xml:"Cd"`

	// Proprietary code to  express the processing status of an advice, a request or a movement.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus3FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus3Code)(&value)
}

func (p *ProcessedStatus3FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
