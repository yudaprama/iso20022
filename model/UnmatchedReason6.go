package model

// The status of an instruction, advice or request.
type UnmatchedReason6 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason9Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason6) AddCode() *UnmatchedReason9Choice {
	u.Code = new(UnmatchedReason9Choice)
	return u.Code
}

func (u *UnmatchedReason6) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
