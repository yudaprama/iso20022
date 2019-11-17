package model

// Status of an instruction, advice or request.
type UnmatchedReason11 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason14Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason11) AddCode() *UnmatchedReason14Choice {
	u.Code = new(UnmatchedReason14Choice)
	return u.Code
}

func (u *UnmatchedReason11) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
