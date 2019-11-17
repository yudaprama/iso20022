package model

// Choice of format for the rejection reason.
type RejectionReason26Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason38Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason26Choice) SetCode(value string) {
	r.Code = (*RejectionReason38Code)(&value)
}

func (r *RejectionReason26Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
