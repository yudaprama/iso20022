package model

// Choice of formats to  express the reason of a rejection of an information advice.
type RejectionReason15FormatChoice struct {

	// Standard code to specify the reason of a rejection of an information advice.
	Code *RejectionReason15Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of an information advice.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason15FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason15Code)(&value)
}

func (r *RejectionReason15FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
