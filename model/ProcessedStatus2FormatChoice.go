package model

// Choice of formats to  express the processing status of a cancellation request.
type ProcessedStatus2FormatChoice struct {

	// Standard code to specify  the processing status of a cancellation request.
	Code *ProcessedStatus2Code `xml:"Cd"`

	// Proprietary code to  express the processing status of a cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus2FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus2Code)(&value)
}

func (p *ProcessedStatus2FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
