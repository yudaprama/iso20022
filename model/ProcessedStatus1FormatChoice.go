package model

// Choice of formats to  express the processing status of the notification advice.
type ProcessedStatus1FormatChoice struct {

	// Standard code to  specify to  express the processing status of the notification advice.
	Code *ProcessedStatus1Code `xml:"Cd"`

	// Proprietary code to  express the processing status of the notification advice.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus1FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus1Code)(&value)
}

func (p *ProcessedStatus1FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
