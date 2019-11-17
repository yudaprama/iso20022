package model

// Choice of formats to  express the reason of a rejection of a standing instruction cancellation request.
type RejectionReason10FormatChoice struct {

	// Standard code to specify the reason of a rejection of a standing instruction cancellation request.
	Code *RejectionReason10Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of a standing instruction cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason10FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason10Code)(&value)
}

func (r *RejectionReason10FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
