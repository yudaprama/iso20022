package model

// Choice of format for the processing status.
type ProcessingStatus48Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus11Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus21Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus12Choice `xml:"Rpr"`

	// Instruction/Request has been cancelled.
	Cancelled *CancellationStatus14Choice `xml:"Canc"`

	// Cancellation request from yourself for this instruction is pending, waiting for further processing.
	PendingCancellation *PendingStatus38Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or consent.
	CancellationRequested *ProprietaryReason4 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason4 `xml:"ModReqd"`
}

func (p *ProcessingStatus48Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus48Choice) AddPendingProcessing() *PendingProcessingStatus11Choice {
	p.PendingProcessing = new(PendingProcessingStatus11Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus48Choice) AddRejected() *RejectionStatus21Choice {
	p.Rejected = new(RejectionStatus21Choice)
	return p.Rejected
}

func (p *ProcessingStatus48Choice) AddRepair() *RepairStatus12Choice {
	p.Repair = new(RepairStatus12Choice)
	return p.Repair
}

func (p *ProcessingStatus48Choice) AddCancelled() *CancellationStatus14Choice {
	p.Cancelled = new(CancellationStatus14Choice)
	return p.Cancelled
}

func (p *ProcessingStatus48Choice) AddPendingCancellation() *PendingStatus38Choice {
	p.PendingCancellation = new(PendingStatus38Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus48Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}

func (p *ProcessingStatus48Choice) AddCancellationRequested() *ProprietaryReason4 {
	p.CancellationRequested = new(ProprietaryReason4)
	return p.CancellationRequested
}

func (p *ProcessingStatus48Choice) AddModificationRequested() *ProprietaryReason4 {
	p.ModificationRequested = new(ProprietaryReason4)
	return p.ModificationRequested
}
