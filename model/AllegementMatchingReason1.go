package model

// The status of an instruction, advice or request.
type AllegementMatchingReason1 struct {

	// Specifies the reason why the instruction has been alleged.
	Code *AllegementReason1Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AllegementMatchingReason1) AddCode() *AllegementReason1Choice {
	a.Code = new(AllegementReason1Choice)
	return a.Code
}

func (a *AllegementMatchingReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
