package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason10 struct {

	// Specifies the reason why the instruction or instruction cancellation has been accepted.
	ReasonCode *AcceptedReason11Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason10) AddReasonCode() *AcceptedReason11Choice {
	a.ReasonCode = new(AcceptedReason11Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason10) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
