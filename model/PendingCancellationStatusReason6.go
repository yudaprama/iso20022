package model

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason6 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason6) AddReasonCode() *PendingCancellationReason4Choice {
	p.ReasonCode = new(PendingCancellationReason4Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason6) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
