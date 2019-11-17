package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason2Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason24Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason2Choice) SetCode(value string) {
	r.Code = (*RejectionReason24Code)(&value)
}

func (r *RejectionAndRepairReason2Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
