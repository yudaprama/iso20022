package model

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason9 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason11Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason9) AddCode() *PendingProcessingReason11Choice {
	p.Code = new(PendingProcessingReason11Choice)
	return p.Code
}

func (p *PendingProcessingReason9) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
