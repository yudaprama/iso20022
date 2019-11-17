package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason17 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason30Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason17) AddCode() *PendingReason30Choice {
	p.Code = new(PendingReason30Choice)
	return p.Code
}

func (p *PendingReason17) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
