package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason13Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason27Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason13Choice) SetCode(value string) {
	r.Code = (*RejectionReason27Code)(&value)
}

func (r *RejectionAndRepairReason13Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
