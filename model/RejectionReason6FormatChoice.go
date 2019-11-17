package model

// Choice of formats to  express the reason of a rejection of the notification advice.
type RejectionReason6FormatChoice struct {

	// Standard code to specify the reason of a rejection of the notification advice.
	Code *RejectionReason6Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of the notification advice.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason6FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason6Code)(&value)
}

func (r *RejectionReason6FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
