package model

// Specifies reasons for the pending status.
type PendingStatusReason13 struct {

	// Specifies the reason why the instruction's processing is pending.
	ReasonCode *PendingReason48Choice `xml:"RsnCd"`

	// Provides additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingStatusReason13) AddReasonCode() *PendingReason48Choice {
	p.ReasonCode = new(PendingReason48Choice)
	return p.ReasonCode
}

func (p *PendingStatusReason13) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
