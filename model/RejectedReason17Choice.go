package model

// Choice of formats for the specification of the rejected reason.
type RejectedReason17Choice struct {

	// Rejected reason expressed as a code.
	Code *CancellationRejectedReason1Code `xml:"Cd"`

	// Rejected reason expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *RejectedReason17Choice) SetCode(value string) {
	r.Code = (*CancellationRejectedReason1Code)(&value)
}

func (r *RejectedReason17Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}
