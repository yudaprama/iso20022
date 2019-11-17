package model

// The status of an instruction, advice or request.
type RepairReason5 struct {

	// Specifies the reason why the instruction is in repair.
	Code *RepairReason9Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason5) AddCode() *RepairReason9Choice {
	r.Code = new(RepairReason9Choice)
	return r.Code
}

func (r *RepairReason5) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
