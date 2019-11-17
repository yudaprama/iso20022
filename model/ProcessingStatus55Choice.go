package model

// Choice of format for the processing status.
type ProcessingStatus55Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus24Choice `xml:"AckdAccptd"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus32Choice `xml:"Rjctd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (p *ProcessingStatus55Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus24Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus24Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus55Choice) AddRejected() *RejectionOrRepairStatus32Choice {
	p.Rejected = new(RejectionOrRepairStatus32Choice)
	return p.Rejected
}

func (p *ProcessingStatus55Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}
