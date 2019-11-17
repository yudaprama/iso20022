package model

// Specifies reasons for the accepted status.
type AcceptedStatusReason11 struct {

	// Specifies the reason why the instruction has been accepted.
	ReasonCode *AcceptedReason12Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (a *AcceptedStatusReason11) AddReasonCode() *AcceptedReason12Choice {
	a.ReasonCode = new(AcceptedReason12Choice)
	return a.ReasonCode
}

func (a *AcceptedStatusReason11) SetAdditionalReasonInformation(value string) {
	a.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
