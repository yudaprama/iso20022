package model

// Choice of format for the rejection reason.
type RejectionReason4Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason25Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason4Choice) SetCode(value string) {
	r.Code = (*RejectionReason25Code)(&value)
}

func (r *RejectionReason4Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
