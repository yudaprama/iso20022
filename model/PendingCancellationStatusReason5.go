package model

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason5 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason3Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason5) AddReasonCode() *PendingCancellationReason3Choice {
	p.ReasonCode = new(PendingCancellationReason3Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
