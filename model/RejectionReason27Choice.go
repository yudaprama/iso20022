package model

// Choice of format for the rejection reason.
type RejectionReason27Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason37Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason27Choice) SetCode(value string) {
	r.Code = (*RejectionReason37Code)(&value)
}

func (r *RejectionReason27Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
