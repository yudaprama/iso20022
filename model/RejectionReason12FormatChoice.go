package model

// Choice of formats to  express the reason of a rejection of a deactivation instruction.
type RejectionReason12FormatChoice struct {

	// Standard code to specify the reason of a rejection of a deactivation instruction.
	Code *RejectionReason12Code `xml:"Cd"`

	// Proprietary code to  express the reason of a rejection of a deactivation instruction.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (r *RejectionReason12FormatChoice) SetCode(value string) {
	r.Code = (*RejectionReason12Code)(&value)
}

func (r *RejectionReason12FormatChoice) AddProprietary() *GenericIdentification13 {
	r.Proprietary = new(GenericIdentification13)
	return r.Proprietary
}
