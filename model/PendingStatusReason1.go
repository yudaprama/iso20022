package model

// Specifies reasons for the pending status.
type PendingStatusReason1 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason6Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason1) AddReasonCode() *PendingReason6Choice {
	p.ReasonCode = new(PendingReason6Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
