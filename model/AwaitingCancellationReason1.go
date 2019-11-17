package model

// The status of an instruction, advice or request.
type AwaitingCancellationReason1 struct {

	// Specifies the reason why the trade is wainting the cancellation.
	Code *AwaitingCancellationReason1Choice `xml:"Cd"`

	// Provides additional information about the reason in narrative form.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AwaitingCancellationReason1) AddCode() *AwaitingCancellationReason1Choice {
	a.Code = new(AwaitingCancellationReason1Choice)
	return a.Code
}

func (a *AwaitingCancellationReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
