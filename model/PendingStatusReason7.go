package model

// Specifies reasons for the pending status.
type PendingStatusReason7 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason24Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason7) AddReasonCode() *PendingReason24Choice {
	p.ReasonCode = new(PendingReason24Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
