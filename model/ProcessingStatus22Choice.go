package model

// Choice of format for the processing status.
type ProcessingStatus22Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus12Choice `xml:"AckdAccptd"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus18Choice `xml:"Rjctd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (p *ProcessingStatus22Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus12Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus12Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus22Choice) AddRejected() *RejectionOrRepairStatus18Choice {
	p.Rejected = new(RejectionOrRepairStatus18Choice)
	return p.Rejected
}

func (p *ProcessingStatus22Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}
