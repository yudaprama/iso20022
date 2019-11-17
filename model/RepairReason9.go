package model

// Specifies the reason why the instruction is in repair.
type RepairReason9 struct {

	// Specifies the reason why the instruction/request has a repair status.
	Code *RepairReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason9) AddCode() *RepairReason10Choice {
	r.Code = new(RepairReason10Choice)
	return r.Code
}

func (r *RepairReason9) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
