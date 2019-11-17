package model

// Specifies the reason why the instruction is in repair.
type RepairReason11 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason13Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason11) AddCode() *RepairReason13Choice {
	r.Code = new(RepairReason13Choice)
	return r.Code
}

func (r *RepairReason11) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
