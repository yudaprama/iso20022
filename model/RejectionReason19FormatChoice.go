package model

// Choice of formats to  express the reason of a rejection of a global movement authorisation request.
type RejectionReason19FormatChoice struct {

	// Standard code to specify the reason of a rejection of a global movement authorisation request.
	Code *RejectionReason19Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of a global movement authorisation request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason19FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason19Code)(&value)
}

func (r *RejectionReason19FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
