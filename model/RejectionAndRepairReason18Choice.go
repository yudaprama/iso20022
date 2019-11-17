package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason18Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason32Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason18Choice) SetCode(value string) {
	r.Code = (*RejectionReason32Code)(&value)
}

func (r *RejectionAndRepairReason18Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
