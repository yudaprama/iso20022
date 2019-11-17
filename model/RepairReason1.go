package model

// The status of an instruction, advice or request.
type RepairReason1 struct {

	// Specifies the reason why the instruction/request has a rejected or repair status.
	Code *RepairReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (r *RepairReason1) AddCode() *RepairReason1Choice {
	r.Code = new(RepairReason1Choice)
	return r.Code
}

func (r *RepairReason1) SetAdditionalReasonInformation(value string) {
	r.AdditionalReasonInformation = (*Max210Text)(&value)
}
