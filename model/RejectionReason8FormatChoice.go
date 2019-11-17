package model

// Choice of formats to  express the reason of a rejection of an election amendment request.
type RejectionReason8FormatChoice struct {

	// Standard code to specify the reason of a rejection of an election amendment request.
	Code *RejectionReason8Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of an election amendment request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason8FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason8Code)(&value)
}

func (r *RejectionReason8FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
