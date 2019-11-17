package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason3 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason3) AddReasonCode() *AcceptedReason3Choice {
	a.ReasonCode = new(AcceptedReason3Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason3) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*Max210Text)(&value)
}
