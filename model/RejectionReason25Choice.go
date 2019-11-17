package model

// Choice of format for the rejection reason.
type RejectionReason25Choice struct {

	// Specifies the reason why the instruction/request has a rejection status using a code.
	Code *RejectionReason40Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a rejection status using a proprietary format.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason25Choice) SetCode(value string) {
	r.Code = (*RejectionReason40Code)(&value)
}

func (r *RejectionReason25Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
