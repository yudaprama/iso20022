package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason23Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason29Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason23Choice) SetCode(value string) {
	r.Code = (*RejectionReason29Code)(&value)
}

func (r *RejectionAndRepairReason23Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
