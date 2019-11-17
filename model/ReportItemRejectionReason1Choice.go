package model

// Provides the report item rejection reason in a coded or proprietary  form.
type ReportItemRejectionReason1Choice struct {

	// Status reason expressed as a code.
	Code *HoldingRejectionReason41Code `xml:"Cd"`

	// Status reason expressed as a proprietary code
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (r *ReportItemRejectionReason1Choice) SetCode(value string) {
	r.Code = (*HoldingRejectionReason41Code)(&value)
}

func (r *ReportItemRejectionReason1Choice) AddProprietary() *GenericIdentification30 {
	r.Proprietary = new(GenericIdentification30)
	return r.Proprietary
}
