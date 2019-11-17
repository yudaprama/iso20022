package model

// The status of an instruction, advice or request.
type AllegmentMatchingReason1 struct {

	// Specifies the reason why the instruction has been alleged.
	Code *AllegmentReason1Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AllegmentMatchingReason1) AddCode() *AllegmentReason1Choice {
	a.Code = new(AllegmentReason1Choice)
	return a.Code
}

func (a *AllegmentMatchingReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
