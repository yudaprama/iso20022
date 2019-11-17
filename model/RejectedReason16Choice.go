package model

// Choice of formats for the specification of the rejected reason.
type RejectedReason16Choice struct {

	// Rejected reason expressed as a code.
	Code *RejectedStatusReason6Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *RejectedReason16Choice) SetCode(value string) {
	r.Code = (*RejectedStatusReason6Code)(&value)
}

func (r *RejectedReason16Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}
