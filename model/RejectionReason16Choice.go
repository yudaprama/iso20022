package model

// Choice of format for the rejection reason.
type RejectionReason16Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason39Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason16Choice) SetCode(value string) {
	r.Code = (*RejectionReason39Code)(&value)
}

func (r *RejectionReason16Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
