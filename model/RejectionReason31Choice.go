package model

// Choice of format for the rejection reason.
type RejectionReason31Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason31Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason31Choice) SetCode(value string) {
	r.Code = (*RejectionReason31Code)(&value)
}

func (r *RejectionReason31Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
