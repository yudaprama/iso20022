package model

// Choice of formats for the specification of the rejected reason.
type RejectedReason4Choice struct {

	// Rejected reason expressed as a code.
	Code *RejectedStatusReason6Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (r *RejectedReason4Choice) SetCode(value string) {
	r.Code = (*RejectedStatusReason6Code)(&value)
}

func (r *RejectedReason4Choice) AddProprietary() *GenericIdentification29 {
	r.Proprietary = new(GenericIdentification29)
	return r.Proprietary
}
