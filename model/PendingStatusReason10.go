package model

// Specifies reasons for the pending status.
type PendingStatusReason10 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason33Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason10) AddReasonCode() *PendingReason33Choice {
	p.ReasonCode = new(PendingReason33Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason10) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
