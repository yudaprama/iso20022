package model

// Choice of formats to  express the processing status of a request.
type ProcessedStatus5FormatChoice struct {

	// Standard code to specify  the processing status of a request.
	Code *ProcessedStatus5Code `xml:"Cd"`

	// Proprietary code to  express the processing status of a request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *ProcessedStatus5FormatChoice) SetCode(value string) {
	p.Code = (*ProcessedStatus5Code)(&value)
}

func (p *ProcessedStatus5FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
