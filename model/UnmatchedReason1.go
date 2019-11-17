package model

// The status of an instruction, advice or request.
type UnmatchedReason1 struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (u *UnmatchedReason1) AddCode() *UnmatchedReason1Choice {
	u.Code = new(UnmatchedReason1Choice)
	return u.Code
}

func (u *UnmatchedReason1) SetAdditionalReasonInformation(value string) {
	u.AdditionalReasonInformation = (*Max210Text)(&value)
}
