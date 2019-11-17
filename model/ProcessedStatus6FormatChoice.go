package model

// Choice of formats to  express the processing status of a deactivation instruction.
type ProcessedStatus6FormatChoice struct {

	// Standard code to specify the processing status of a deactivation instruction.
	Code *ProcessedStatus6Code `xml:"Cd"`

	// Proprietary code to  express the processing status of a deactivation instruction.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus6FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus6Code)(&value)
}

func (p *ProcessedStatus6FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
