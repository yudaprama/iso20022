package model

// Choice of formats to  express the processing status.
type ProcessingStatus1FormatChoice struct {

	// Standard code to specify the processing status.
	Code *ProcessingStatus1Code `xml:"Cd"`

	// Proprietary code to  express the processing status.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessingStatus1FormatChoice) SetCode(value string) {
	p.Code = (*ProcessingStatus1Code)(&value)
}

func (p *ProcessingStatus1FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
