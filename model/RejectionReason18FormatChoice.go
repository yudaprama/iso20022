package model

// Choice of formats to  express the reason of a rejection of an election advice.
type RejectionReason18FormatChoice struct {

	// Standard code to specify the reason of a rejection of an election advice.
	Code *RejectionReason18Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of an election advice.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason18FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason18Code)(&value)
}

func (r *RejectionReason18FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
