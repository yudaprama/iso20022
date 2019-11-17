package model

// Reason for a pending status.
type PendingProcessingStatusReason1 struct {

	// Reason for the pending status.
	Reason *PendingProcessingReason9Choice `xml:"Rsn"`

	// Additional information about the processed instruction.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (p *PendingProcessingStatusReason1) AddReason() *PendingProcessingReason9Choice {
	p.Reason = new(PendingProcessingReason9Choice)
	return p.Reason
}

func (p *PendingProcessingStatusReason1) SetAdditionalReasonInformation(value string) {
	p.AdditionalReasonInformation = (*Max210Text)(&value)
}
