package model

// Specifies reasons for the pending status.
type PendingStatusReason5 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason22Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason5) AddReasonCode() *PendingReason22Choice {
	p.ReasonCode = new(PendingReason22Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason5) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
