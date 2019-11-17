package model

// Choice of formats to  express the processing status of the standing instruction cancellation request.
type ProcessedStatus4FormatChoice struct {

	// Standard code to specify  the processing status of the standing instruction cancellation request.
	Code *ProcessedStatus4Code `xml:"Cd"`

	// Proprietary code to  express the processing status of the standing instruction cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus4FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus4Code)(&value)
}

func (p *ProcessedStatus4FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
