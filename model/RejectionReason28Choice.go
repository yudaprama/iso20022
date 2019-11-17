package model

// Choice of format for the rejection reason.
type RejectionReason28Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason38Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason28Choice) SetCode(value string) {
	r.Code = (*RejectionReason38Code)(&value)
}

func (r *RejectionReason28Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
