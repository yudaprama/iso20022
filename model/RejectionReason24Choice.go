package model

// Choice of format for the rejection reason.
type RejectionReason24Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason31Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason24Choice) SetCode(value string) {
	r.Code = (*RejectionReason31Code)(&value)
}

func (r *RejectionReason24Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
