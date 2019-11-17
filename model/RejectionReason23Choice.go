package model

// Choice of format for the rejection reason.
type RejectionReason23Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason30Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionReason23Choice) SetCode(value string) {
	r.Code = (*RejectionReason30Code)(&value)
}

func (r *RejectionReason23Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
