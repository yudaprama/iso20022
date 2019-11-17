package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a rejected status.
type RejectedReason9Choice struct {

	// Standard code to specify the reason why the instruction/cancellation request has a rejected status.
	Code *RejectionReason43Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/cancellation request has a rejected status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectedReason9Choice) SetCode(value string) {
	r.Code = (*RejectionReason43Code)(&value)
}

func (r *RejectedReason9Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
