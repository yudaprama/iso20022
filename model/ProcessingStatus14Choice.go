package model

// Choice of format for the processing status.
type ProcessingStatus14Choice struct {

	// Request has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Modification Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus11Choice `xml:"Rjctd"`

	// Modification request was completed.
	Completed *NoSpecifiedReason1 `xml:"Cmpltd"`

	// Modification request will not be executed.
	Denied *DeniedStatus2Choice `xml:"Dnd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Modification is pending. It is not known at this time whether modification can be affected.
	Pending *PendingStatus4Choice `xml:"Pdg"`
}

func (p *ProcessingStatus14Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus14Choice) AddRejected() *RejectionOrRepairStatus11Choice {
	p.Rejected = new(RejectionOrRepairStatus11Choice)
	return p.Rejected
}

func (p *ProcessingStatus14Choice) AddCompleted() *NoSpecifiedReason1 {
	p.Completed = new(NoSpecifiedReason1)
	return p.Completed
}

func (p *ProcessingStatus14Choice) AddDenied() *DeniedStatus2Choice {
	p.Denied = new(DeniedStatus2Choice)
	return p.Denied
}

func (p *ProcessingStatus14Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus14Choice) AddPending() *PendingStatus4Choice {
	p.Pending = new(PendingStatus4Choice)
	return p.Pending
}
