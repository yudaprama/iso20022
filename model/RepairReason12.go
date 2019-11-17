package model

// Specifies the reason why the instruction is in repair.
type RepairReason12 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason12) AddCode() *RepairReason14Choice {
	r.Code = new(RepairReason14Choice)
	return r.Code
}

func (r *RepairReason12) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
