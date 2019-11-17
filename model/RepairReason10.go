package model

// Specifies the reason why the instruction is in repair.
type RepairReason10 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason10) AddCode() *RepairReason12Choice {
	r.Code = new(RepairReason12Choice)
	return r.Code
}

func (r *RepairReason10) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
