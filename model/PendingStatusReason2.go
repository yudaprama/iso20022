package model

// Specifies reasons for the pending status.
type PendingStatusReason2 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason4Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason2) AddReasonCode() *PendingReason4Choice {
	p.ReasonCode = new(PendingReason4Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason2) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
