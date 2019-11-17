package model

// Specifies the reason why the instruction or request has a pending status.
type PendingReason18 struct {

	// Specifies the reason why the instruction has a pending status.
	Code *PendingReason31Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingReason18) AddCode() *PendingReason31Choice {
	p.Code = new(PendingReason31Choice)
	return p.Code
}

func (p *PendingReason18) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
