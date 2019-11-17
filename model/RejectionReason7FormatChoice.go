package model

// Choice of formats to  express the reason of a rejection cancellation request.
type RejectionReason7FormatChoice struct {

	// Standard code to specify the reason of a rejection cancellation request.
	Code *RejectionReason7Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason7FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason7Code)(&value)
}

func (r *RejectionReason7FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
