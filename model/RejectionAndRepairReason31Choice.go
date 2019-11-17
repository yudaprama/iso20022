package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason31Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason24Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason31Choice) SetCode(value string) {
	r.Code = (*RejectionReason24Code)(&value)
}

func (r *RejectionAndRepairReason31Choice) AddProprietary() *GenericIdentification47 {
	r.Proprietary = new(GenericIdentification47)
	return r.Proprietary
}
