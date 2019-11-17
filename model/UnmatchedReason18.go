package model

// Status of an instruction, advice or request.
type UnmatchedReason18 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason25Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason18) AddCode() *UnmatchedReason25Choice {
	u.Code = new(UnmatchedReason25Choice)
	return u.Code
}

func (u *UnmatchedReason18) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
