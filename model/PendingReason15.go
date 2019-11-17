package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason15 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason27Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason15) AddCode() *PendingReason27Choice {
	p.Code = new(PendingReason27Choice)
	return p.Code
}

func (p *PendingReason15) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
