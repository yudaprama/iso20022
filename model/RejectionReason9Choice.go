package model

// Choice of format for the rejection reason.
type RejectionReason9Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason28Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (r *RejectionReason9Choice) SetCode(value string) {
	r.Code = (*RejectionReason28Code)(&value)
}

func (r *RejectionReason9Choice) AddProprietary() *GenericIdentification38 {
	r.Proprietary = new(GenericIdentification38)
	return r.Proprietary
}
