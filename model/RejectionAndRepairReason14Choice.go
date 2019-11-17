package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason14Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason29Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason14Choice) SetCode(value string) {
	r.Code = (*RejectionReason29Code)(&value)
}

func (r *RejectionAndRepairReason14Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
