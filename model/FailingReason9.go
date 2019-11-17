package model

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason9 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason9) AddCode() *FailingReason10Choice {
	f.Code = new(FailingReason10Choice)
	return f.Code
}

func (f *FailingReason9) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
