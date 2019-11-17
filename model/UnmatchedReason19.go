package model

// Specifies the reason why the instruction has an unmatched status.
type UnmatchedReason19 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason26Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason19) AddCode() *UnmatchedReason26Choice {
	u.Code = new(UnmatchedReason26Choice)
	return u.Code
}

func (u *UnmatchedReason19) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
