package model

// The status of an instruction, advice or request.
type AwaitingAffirmationReason1 struct {

	// Specifies the reason why the trade is wainting the affirmation.
	Code *AwaitingAffirmationReason1Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AwaitingAffirmationReason1) AddCode() *AwaitingAffirmationReason1Choice {
	a.Code = new(AwaitingAffirmationReason1Choice)
	return a.Code
}

func (a *AwaitingAffirmationReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
