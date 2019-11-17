package model

// Choice of formats to  express the reason of a rejection of the standing instruction request.
type RejectionReason20FormatChoice struct {

	// Standard code to specify the reason of a rejection of the standing instruction request.
	Code *RejectionReason20Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of the standing instruction request.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason20FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason20Code)(&value)
}

func (r *RejectionReason20FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
