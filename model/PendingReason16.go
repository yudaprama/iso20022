package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason16 struct {

	// Specifies the reason why a cancellation request sent for the related instruction is pending.
	Code *PendingReason28Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason16) AddCode() *PendingReason28Choice {
	p.Code = new(PendingReason28Choice)
	return p.Code
}

func (p *PendingReason16) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
