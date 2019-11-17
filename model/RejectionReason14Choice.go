package model

// Choice of format for the rejection reason.
type RejectionReason14Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason37Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason14Choice) SetCode(value string) {
	r.Code = (*RejectionReason37Code)(&value)
}

func (r *RejectionReason14Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
