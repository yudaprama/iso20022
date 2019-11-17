package model

// Choice of format for the processing status.
type ProcessingStatus61Choice struct {

	// The cancellation is pending processing.
	PendingCancellation *PendingStatus51Choice `xml:"PdgCxl"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus35Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus34Choice `xml:"Rpr"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus31Choice `xml:"AckdAccptd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus21Choice `xml:"Dnd"`

	// Cancellation requested executed.
	Cancelled *CancellationStatus20Choice `xml:"Canc"`
}

func (p *ProcessingStatus61Choice) AddPendingCancellation() *PendingStatus51Choice {
	p.PendingCancellation = new(PendingStatus51Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus61Choice) AddRejected() *RejectionOrRepairStatus35Choice {
	p.Rejected = new(RejectionOrRepairStatus35Choice)
	return p.Rejected
}

func (p *ProcessingStatus61Choice) AddRepair() *RejectionOrRepairStatus34Choice {
	p.Repair = new(RejectionOrRepairStatus34Choice)
	return p.Repair
}

func (p *ProcessingStatus61Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus31Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus31Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus61Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}

func (p *ProcessingStatus61Choice) AddDenied() *DeniedStatus21Choice {
	p.Denied = new(DeniedStatus21Choice)
	return p.Denied
}

func (p *ProcessingStatus61Choice) AddCancelled() *CancellationStatus20Choice {
	p.Cancelled = new(CancellationStatus20Choice)
	return p.Cancelled
}
