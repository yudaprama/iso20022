package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a rejected status.
type RejectedReason10Choice struct {

	// Standard code to specify the reason why the instruction/cancellation request has a rejected status.
	Code *RejectionReason45Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/cancellation request has a rejected status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectedReason10Choice) SetCode(value string) {
	r.Code = (*RejectionReason45Code)(&value)
}

func (r *RejectedReason10Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
