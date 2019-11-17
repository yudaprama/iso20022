package model

// Choice of format for the processing status.
type ProcessingStatus20Choice struct {

	// The cancellation is pending processing.
	PendingCancellation *PendingStatus11Choice `xml:"PdgCxl"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus25Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus14Choice `xml:"Rpr"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus12Choice `xml:"AckdAccptd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus6Choice `xml:"Dnd"`

	// Cancellation requested executed.
	Cancelled *CancellationStatus9Choice `xml:"Canc"`
}

func (p *ProcessingStatus20Choice) AddPendingCancellation() *PendingStatus11Choice {
	p.PendingCancellation = new(PendingStatus11Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus20Choice) AddRejected() *RejectionOrRepairStatus25Choice {
	p.Rejected = new(RejectionOrRepairStatus25Choice)
	return p.Rejected
}

func (p *ProcessingStatus20Choice) AddRepair() *RejectionOrRepairStatus14Choice {
	p.Repair = new(RejectionOrRepairStatus14Choice)
	return p.Repair
}

func (p *ProcessingStatus20Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus12Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus12Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus20Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus20Choice) AddDenied() *DeniedStatus6Choice {
	p.Denied = new(DeniedStatus6Choice)
	return p.Denied
}

func (p *ProcessingStatus20Choice) AddCancelled() *CancellationStatus9Choice {
	p.Cancelled = new(CancellationStatus9Choice)
	return p.Cancelled
}
