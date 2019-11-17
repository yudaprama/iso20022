package model

// Choice of format for the rejection or repair reason.
type RejectionAndRepairReason1Choice struct {

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Code *RejectionReason21Code `xml:"Cd"`

	// Specifies the reason why the instruction/request has a repair or rejection status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (r *RejectionAndRepairReason1Choice) SetCode(value string) {
	r.Code = (*RejectionReason21Code)(&value)
}

func (r *RejectionAndRepairReason1Choice) AddProprietary() *GenericIdentification20 {
	r.Proprietary = new(GenericIdentification20)
	return r.Proprietary
}
