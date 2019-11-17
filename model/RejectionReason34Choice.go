package model

// Choice of format for the rejection reason.
type RejectionReason34Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason39Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason34Choice) SetCode(value string) {
	r.Code = (*RejectionReason39Code)(&value)
}

func (r *RejectionReason34Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
