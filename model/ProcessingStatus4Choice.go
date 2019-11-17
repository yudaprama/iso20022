package model

// Choice of format for the processing status.
type ProcessingStatus4Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus1Choice `xml:"AckdAccptd"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus5Choice `xml:"Rjctd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`
}

func (p *ProcessingStatus4Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus1Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus1Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus4Choice) AddRejected() *RejectionOrRepairStatus5Choice {
	p.Rejected = new(RejectionOrRepairStatus5Choice)
	return p.Rejected
}

func (p *ProcessingStatus4Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}
