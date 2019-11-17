package model

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason8 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason6Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason8) AddReasonCode() *PendingCancellationReason6Choice {
	p.ReasonCode = new(PendingCancellationReason6Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
