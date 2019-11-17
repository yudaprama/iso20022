package model

// Choice of format for the rejection reason.
type RejectionReason33Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason37Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason33Choice) SetCode(value string) {
	r.Code = (*RejectionReason37Code)(&value)
}

func (r *RejectionReason33Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
