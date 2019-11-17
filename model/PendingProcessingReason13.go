package model

// Specifies the reason why the instruction or request has a pending processing status.
type PendingProcessingReason13 struct {

	// Specifies the reason why the instruction has a pending processing status.
	Code *PendingProcessingReason15Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *RestrictedFINXMax210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingReason13) AddCode() *PendingProcessingReason15Choice {
	p.Code = new(PendingProcessingReason15Choice)
	return p.Code
}

func (p *PendingProcessingReason13) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*RestrictedFINXMax210Text)(&value)
}
