package model

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason7 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason5Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason7) AddReasonCode() *PendingCancellationReason5Choice {
	p.ReasonCode = new(PendingCancellationReason5Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
