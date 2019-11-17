package model

// Choice of formats to  express the reason of a rejection of an election cancellation request.
type RejectionReason9FormatChoice struct {

	// Standard code to specify the reason of a rejection of an election cancellation request.
	Code *RejectionReason9Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of an election cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason9FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason9Code)(&value)
}

func (r *RejectionReason9FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
