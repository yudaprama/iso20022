package model

// Choice of format for the rejection reason.
type RejectionReason11Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason31Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason11Choice) SetCode(value string) {
	r.Code = (*RejectionReason31Code)(&value)
}

func (r *RejectionReason11Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
