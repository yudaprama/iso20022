package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a rejected status.
type RejectedReason14Choice struct {

	// Standard code to specify the reason why the instruction/cancellation request has a rejected status.
	Code *RejectionReason45Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/cancellation request has a rejected status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectedReason14Choice) SetCode(value string) {
	r.Code = (*RejectionReason45Code)(&value)
}

func (r *RejectedReason14Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
