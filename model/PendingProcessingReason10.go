package model

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason10 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason12Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason10) AddCode() *PendingProcessingReason12Choice {
	p.Code = new(PendingProcessingReason12Choice)
	return p.Code
}

func (p *PendingProcessingReason10) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
