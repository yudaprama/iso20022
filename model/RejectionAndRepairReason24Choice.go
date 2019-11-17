package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason24Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason32Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason24Choice) SetCode(value string) {
	r.Code = (*RejectionReason32Code)(&value)
}

func (r *RejectionAndRepairReason24Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
