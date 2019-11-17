package model

// Status of an instruction, advice or request.
type UnmatchedReason12 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason12) AddCode() *UnmatchedReason15Choice {
	u.Code = new(UnmatchedReason15Choice)
	return u.Code
}

func (u *UnmatchedReason12) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
