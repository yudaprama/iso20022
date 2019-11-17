package model

// Choice of formats to  express the reason of a rejection of the notification cancellation request.
type RejectionReason11FormatChoice struct {

	// Standard code to specify the reason of a rejection of the notification cancellation request.
	Code *RejectionReason11Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of the notification cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason11FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason11Code)(&value)
}

func (r *RejectionReason11FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
