package model

// Specifies reasons for the pending cancellation status.
type PendingCancellationStatusReason3 struct {

	// Specifies the reason why the cancellation request is pending.
	ReasonCode *PendingCancellationReason1Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingCancellationStatusReason3) AddReasonCode() *PendingCancellationReason1Choice {
	p.ReasonCode = new(PendingCancellationReason1Choice)
	return p.ReasonCode
}

func (p *PendingCancellationStatusReason3) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
