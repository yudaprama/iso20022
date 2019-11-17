package model

// Choice of format for the processing status.
type ProcessingStatus45Choice struct {

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
	Denied *DeniedStatus10Choice `xml:"Dnd"`

	// Cancellation requested executed.
	Cancelled *CancellationStatus9Choice `xml:"Canc"`
}

func (p *ProcessingStatus45Choice) AddPendingCancellation() *PendingStatus11Choice {
	p.PendingCancellation = new(PendingStatus11Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus45Choice) AddRejected() *RejectionOrRepairStatus25Choice {
	p.Rejected = new(RejectionOrRepairStatus25Choice)
	return p.Rejected
}

func (p *ProcessingStatus45Choice) AddRepair() *RejectionOrRepairStatus14Choice {
	p.Repair = new(RejectionOrRepairStatus14Choice)
	return p.Repair
}

func (p *ProcessingStatus45Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus12Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus12Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus45Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus45Choice) AddDenied() *DeniedStatus10Choice {
	p.Denied = new(DeniedStatus10Choice)
	return p.Denied
}

func (p *ProcessingStatus45Choice) AddCancelled() *CancellationStatus9Choice {
	p.Cancelled = new(CancellationStatus9Choice)
	return p.Cancelled
}
