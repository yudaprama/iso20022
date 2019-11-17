package model

// Specifies the reason of the PendingStatus.
type PendingReason2 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason2Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason2) AddCode() *PendingReason2Choice {
	p.Code = new(PendingReason2Choice)
	return p.Code
}

func (p *PendingReason2) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
