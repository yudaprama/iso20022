package model

// Status of an instruction, advice or request.
type UnmatchedReason15 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason21Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason15) AddCode() *UnmatchedReason21Choice {
	u.Code = new(UnmatchedReason21Choice)
	return u.Code
}

func (u *UnmatchedReason15) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
