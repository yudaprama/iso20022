package model

// The status of an instruction, advice or request.
type UnmatchedReason2 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason2) AddCode() *UnmatchedReason3Choice {
	u.Code = new(UnmatchedReason3Choice)
	return u.Code
}

func (u *UnmatchedReason2) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
