package model

// Choice of format for the processing status.
type ProcessingStatus58Choice struct {

	// Request has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Modification Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus34Choice `xml:"Rjctd"`

	// Modification request was completed.
	Completed *ProprietaryReason5 `xml:"Cmpltd"`

	// Modification request will not be executed.
	Denied *DeniedStatus19Choice `xml:"Dnd"`

	// Modification is pending. It is not known at this time whether modification can be affected.
	Pending *PendingStatus46Choice `xml:"Pdg"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`
}

func (p *ProcessingStatus58Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus58Choice) AddRejected() *RejectionOrRepairStatus34Choice {
	p.Rejected = new(RejectionOrRepairStatus34Choice)
	return p.Rejected
}

func (p *ProcessingStatus58Choice) AddCompleted() *ProprietaryReason5 {
	p.Completed = new(ProprietaryReason5)
	return p.Completed
}

func (p *ProcessingStatus58Choice) AddDenied() *DeniedStatus19Choice {
	p.Denied = new(DeniedStatus19Choice)
	return p.Denied
}

func (p *ProcessingStatus58Choice) AddPending() *PendingStatus46Choice {
	p.Pending = new(PendingStatus46Choice)
	return p.Pending
}

func (p *ProcessingStatus58Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}
