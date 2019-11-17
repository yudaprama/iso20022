package model

// Specifies reasons for the pending status.
type PendingStatusReason12 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason35Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason12) AddReasonCode() *PendingReason35Choice {
	p.ReasonCode = new(PendingReason35Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason12) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
