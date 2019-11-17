package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason1 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason1) AddReasonCode() *AcceptedReason1Choice {
	a.ReasonCode = new(AcceptedReason1Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason1) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
