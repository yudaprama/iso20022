package model

// Specifies the reason why the instruction has an unmatched status.
type UnmatchedReason16 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason23Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason16) AddCode() *UnmatchedReason23Choice {
	u.Code = new(UnmatchedReason23Choice)
	return u.Code
}

func (u *UnmatchedReason16) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
