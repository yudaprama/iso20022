package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason10Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason21Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason10Choice) SetCode(value string) {
	r.Code = (*RejectionReason21Code)(&value)
}

func (r *RejectionAndRepairReason10Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
