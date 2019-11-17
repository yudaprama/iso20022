package model

// The status of an instruction, advice or request.
type FailingReason3 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason3Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason3) AddCode() *FailingReason3Choice {
	f.Code = new(FailingReason3Choice)
	return f.Code
}

func (f *FailingReason3) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}
