package model

// Choice between formats for a rejected reason.
type RejectedReason20Choice struct {

	// Rejected reason expressed as a code.
	Code *RejectedStatusReason11Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (r *RejectedReason20Choice) SetCode(value string) {
	r.Code = (*RejectedStatusReason11Code)(&value)
}

func (r *RejectedReason20Choice) AddProprietary() *GenericIdentification1 {
	r.Proprietary = new(GenericIdentification1)
	return r.Proprietary
}
