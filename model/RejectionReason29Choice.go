package model

// Choice of format for the rejection reason.
type RejectionReason29Choice struct {

	// Specifies the reason why the instruction/request has a rejection status using a code.
	Code *RejectionReason40Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a rejection status using a proprietary format.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason29Choice) SetCode(value string) {
	r.Code = (*RejectionReason40Code)(&value)
}

func (r *RejectionReason29Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
