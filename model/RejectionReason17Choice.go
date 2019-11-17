package model

// Choice of format for the rejection reason.
type RejectionReason17Choice struct {

	// Specifies the reason why the instruction/request has a rejection status using a code.
	Code *RejectionReason40Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a rejection status using a proprietary format.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason17Choice) SetCode(value string) {
	r.Code = (*RejectionReason40Code)(&value)
}

func (r *RejectionReason17Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
