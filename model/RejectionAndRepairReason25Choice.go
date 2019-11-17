package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason25Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason27Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason25Choice) SetCode(value string) {
	r.Code = (*RejectionReason27Code)(&value)
}

func (r *RejectionAndRepairReason25Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
