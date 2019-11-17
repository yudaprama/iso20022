package model

// Choice of format for the processing status.
type ProcessingStatus51Choice struct {

	// Cancellation request from yourself for this instruction is pending waiting for further processing.
	PendingCancellation *PendingStatus38Choice `xml:"PdgCxl"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus12Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus19Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus14Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus16Choice `xml:"Canc"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	CancellationRequested *ProprietaryReason4 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason4 `xml:"ModReqd"`
}

func (p *ProcessingStatus51Choice) AddPendingCancellation() *PendingStatus38Choice {
	p.PendingCancellation = new(PendingStatus38Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus51Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus51Choice) AddPendingProcessing() *PendingProcessingStatus12Choice {
	p.PendingProcessing = new(PendingProcessingStatus12Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus51Choice) AddRejected() *RejectionStatus19Choice {
	p.Rejected = new(RejectionStatus19Choice)
	return p.Rejected
}

func (p *ProcessingStatus51Choice) AddRepair() *RepairStatus14Choice {
	p.Repair = new(RepairStatus14Choice)
	return p.Repair
}

func (p *ProcessingStatus51Choice) AddCancelled() *CancellationStatus16Choice {
	p.Cancelled = new(CancellationStatus16Choice)
	return p.Cancelled
}

func (p *ProcessingStatus51Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}

func (p *ProcessingStatus51Choice) AddCancellationRequested() *ProprietaryReason4 {
	p.CancellationRequested = new(ProprietaryReason4)
	return p.CancellationRequested
}

func (p *ProcessingStatus51Choice) AddModificationRequested() *ProprietaryReason4 {
	p.ModificationRequested = new(ProprietaryReason4)
	return p.ModificationRequested
}
