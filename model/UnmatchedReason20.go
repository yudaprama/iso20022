package model

// Status of an instruction, advice or request.
type UnmatchedReason20 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason20) AddCode() *UnmatchedReason27Choice {
	u.Code = new(UnmatchedReason27Choice)
	return u.Code
}

func (u *UnmatchedReason20) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
