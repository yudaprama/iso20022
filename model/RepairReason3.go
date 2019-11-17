package model

// The status of an instruction, advice or request.
type RepairReason3 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason3) AddCode() *RepairReason2Choice {
	r.Code = new(RepairReason2Choice)
	return r.Code
}

func (r *RepairReason3) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
