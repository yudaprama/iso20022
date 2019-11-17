package model

// Specifies reasons for the pending status.
type PendingStatusReason9 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason32Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason9) AddReasonCode() *PendingReason32Choice {
	p.ReasonCode = new(PendingReason32Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason9) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
