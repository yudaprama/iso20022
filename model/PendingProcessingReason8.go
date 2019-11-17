package model

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason8 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason10Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason8) AddCode() *PendingProcessingReason10Choice {
	p.Code = new(PendingProcessingReason10Choice)
	return p.Code
}

func (p *PendingProcessingReason8) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
