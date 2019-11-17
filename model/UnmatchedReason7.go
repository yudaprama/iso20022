package model

// The status of an instruction, advice or request.
type UnmatchedReason7 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason7) AddCode() *UnmatchedReason10Choice {
	u.Code = new(UnmatchedReason10Choice)
	return u.Code
}

func (u *UnmatchedReason7) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
