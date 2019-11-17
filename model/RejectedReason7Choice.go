package model

// Choice of formats for the reason of a rejected status.
type RejectedReason7Choice struct {

	// Reason for the rejected status expressed as an ISO 20022 code.
	Code *ExternalRejectedReason1Code `xml:"Cd"`

	// Reason for the rejected status expressed as a proprietary code.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (r *RejectedReason7Choice) SetCode(value string) {
	r.Code = (*ExternalRejectedReason1Code)(&value)
}

func (r *RejectedReason7Choice) AddProprietary() *GenericIdentification36 {
	r.Proprietary = new(GenericIdentification36)
	return r.Proprietary
}
