package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason3Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason23Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason3Choice) SetCode(value string) {
	r.Code = (*RejectionReason23Code)(&value)
}

func (r *RejectionAndRepairReason3Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
