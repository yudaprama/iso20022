package model

// Choice between formats for a rejected reason.
type RejectedReason21Choice struct {

	// Rejected reason expressed as a code.
	Code *RejectedStatusReason8Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (r *RejectedReason21Choice) SetCode(value string) {
	r.Code = (*RejectedStatusReason8Code)(&value)
}

func (r *RejectedReason21Choice) AddProprietary() *GenericIdentification1 {
	r.Proprietary = new(GenericIdentification1)
	return r.Proprietary
}
