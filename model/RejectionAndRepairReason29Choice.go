package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason29Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason32Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason29Choice) SetCode(value string) {
	r.Code = (*RejectionReason32Code)(&value)
}

func (r *RejectionAndRepairReason29Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
