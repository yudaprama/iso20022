package model

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason10 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason10) AddCode() *FailingReason11Choice {
	f.Code = new(FailingReason11Choice)
	return f.Code
}

func (f *FailingReason10) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
