package model

// Choice of format for the rejection reason.
type RejectionReason30Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason30Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionReason30Choice) SetCode(value string) {
	r.Code = (*RejectionReason30Code)(&value)
}

func (r *RejectionReason30Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
