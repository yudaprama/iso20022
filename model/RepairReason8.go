package model

// Specifies the reason why the instruction is in repair.
type RepairReason8 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason8) AddCode() *RepairReason10Choice {
	r.Code = new(RepairReason10Choice)
	return r.Code
}

func (r *RepairReason8) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
