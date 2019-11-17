package model

// The status of an instruction, advice or request.
type FailingReason1 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason1Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason1) AddCode() *FailingReason1Choice {
	f.Code = new(FailingReason1Choice)
	return f.Code
}

func (f *FailingReason1) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}
