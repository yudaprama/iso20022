package model

// Specifies the reason of the PendingStatus.
type PendingReason7 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason7) AddCode() *PendingReason15Choice {
	p.Code = new(PendingReason15Choice)
	return p.Code
}

func (p *PendingReason7) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
