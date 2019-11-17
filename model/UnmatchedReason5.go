package model

// The status of an instruction, advice or request.
type UnmatchedReason5 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason7Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason5) AddCode() *UnmatchedReason7Choice {
	u.Code = new(UnmatchedReason7Choice)
	return u.Code
}

func (u *UnmatchedReason5) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
