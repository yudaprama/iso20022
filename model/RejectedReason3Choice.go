package model

// Choice of formats for the specification of the rejected reason.
type RejectedReason3Choice struct {

	// Rejected reason expressed as a code.
	Code *RejectedStatusReason6Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectedReason3Choice) SetCode(value string) {
	r.Code = (*RejectedStatusReason6Code)(&value)
}

func (r *RejectedReason3Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
