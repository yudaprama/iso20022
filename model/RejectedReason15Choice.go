package model

// Choice of formats for the specification of the rejected reason.
type RejectedReason15Choice struct {

	// Rejected reason expressed as a code.
	Code *TransferRejectedStatusReason2Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *RejectedReason15Choice) SetCode(value string) {
	r.Code = (*TransferRejectedStatusReason2Code)(&value)
}

func (r *RejectedReason15Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}
