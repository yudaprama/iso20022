package model

// Status of an instruction, advice or request.
type UnmatchedReason17 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason24Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason17) AddCode() *UnmatchedReason24Choice {
	u.Code = new(UnmatchedReason24Choice)
	return u.Code
}

func (u *UnmatchedReason17) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
