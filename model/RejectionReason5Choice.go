package model

// Choice of format for the rejection reason.
type RejectionReason5Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason16Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionReason5Choice) SetCode(value string) {
	r.Code = (*RejectionReason16Code)(&value)
}

func (r *RejectionReason5Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
