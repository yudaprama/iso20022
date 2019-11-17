package model

// Choice of format for the processing status.
type ProcessingStatus64Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus31Choice `xml:"AckdAccptd"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus37Choice `xml:"Rjctd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (p *ProcessingStatus64Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus31Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus31Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus64Choice) AddRejected() *RejectionOrRepairStatus37Choice {
	p.Rejected = new(RejectionOrRepairStatus37Choice)
	return p.Rejected
}

func (p *ProcessingStatus64Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}
