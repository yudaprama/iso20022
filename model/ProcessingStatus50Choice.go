package model

// Choice of format for the processing status.
type ProcessingStatus50Choice struct {

	// Request has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Modification Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus31Choice `xml:"Rjctd"`

	// Modification request was completed.
	Completed *ProprietaryReason4 `xml:"Cmpltd"`

	// Modification request will not be executed.
	Denied *DeniedStatus15Choice `xml:"Dnd"`

	// Modification is pending. It is not known at this time whether modification can be affected.
	Pending *PendingStatus38Choice `xml:"Pdg"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`
}

func (p *ProcessingStatus50Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus50Choice) AddRejected() *RejectionOrRepairStatus31Choice {
	p.Rejected = new(RejectionOrRepairStatus31Choice)
	return p.Rejected
}

func (p *ProcessingStatus50Choice) AddCompleted() *ProprietaryReason4 {
	p.Completed = new(ProprietaryReason4)
	return p.Completed
}

func (p *ProcessingStatus50Choice) AddDenied() *DeniedStatus15Choice {
	p.Denied = new(DeniedStatus15Choice)
	return p.Denied
}

func (p *ProcessingStatus50Choice) AddPending() *PendingStatus38Choice {
	p.Pending = new(PendingStatus38Choice)
	return p.Pending
}

func (p *ProcessingStatus50Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}
