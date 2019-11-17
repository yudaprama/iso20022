package model

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason8 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason8Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason8) AddCode() *FailingReason8Choice {
	f.Code = new(FailingReason8Choice)
	return f.Code
}

func (f *FailingReason8) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}
