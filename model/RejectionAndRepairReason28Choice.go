package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason28Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason27Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason28Choice) SetCode(value string) {
	r.Code = (*RejectionReason27Code)(&value)
}

func (r *RejectionAndRepairReason28Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
