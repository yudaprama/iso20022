package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason27Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason29Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason27Choice) SetCode(value string) {
	r.Code = (*RejectionReason29Code)(&value)
}

func (r *RejectionAndRepairReason27Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
