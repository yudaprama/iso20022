package model

// Specifies the reason why the instruction or request has a failing settlement status.
type FailingReason7 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason7Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason7) AddCode() *FailingReason7Choice {
	f.Code = new(FailingReason7Choice)
	return f.Code
}

func (f *FailingReason7) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}
