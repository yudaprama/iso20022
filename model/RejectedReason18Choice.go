package model

// Choice between a standard code or proprietary code to specify the reason why the instruction or cancellation request has a rejected status.
type RejectedReason18Choice struct {

	// Standard code to specify the reason why the instruction/cancellation request has a rejected status.
	Code *RejectionReason45Code `xml:"Cd"`

	// Proprietary identification of the reason why the instruction/cancellation request has a rejected status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectedReason18Choice) SetCode(value string) {
	r.Code = (*RejectionReason45Code)(&value)
}

func (r *RejectedReason18Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
