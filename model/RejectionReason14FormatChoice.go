package model

// Choice of formats to  express the reason of a rejection of a movement cancellation request.
type RejectionReason14FormatChoice struct {

	// Standard code to specify the reason of a rejection of a movement cancellation request.
	Code *RejectionReason14Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of a movement cancellation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason14FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason14Code)(&value)
}

func (r *RejectionReason14FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
