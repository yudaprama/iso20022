package model

// Choice of format for the rejection reason.
type RejectionReason22Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason39Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason22Choice) SetCode(value string) {
	r.Code = (*RejectionReason39Code)(&value)
}

func (r *RejectionReason22Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
