package model

// Choice of format for the processing status.
type ProcessingStatus19Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus3Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus5Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus5Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus7Choice `xml:"Canc"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus13Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or consent.
	CancellationRequested *ProprietaryReason1 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason1 `xml:"ModReqd"`
}

func (p *ProcessingStatus19Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus19Choice) AddPendingProcessing() *PendingProcessingStatus3Choice {
	p.PendingProcessing = new(PendingProcessingStatus3Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus19Choice) AddRejected() *RejectionStatus5Choice {
	p.Rejected = new(RejectionStatus5Choice)
	return p.Rejected
}

func (p *ProcessingStatus19Choice) AddRepair() *RepairStatus5Choice {
	p.Repair = new(RepairStatus5Choice)
	return p.Repair
}

func (p *ProcessingStatus19Choice) AddCancelled() *CancellationStatus7Choice {
	p.Cancelled = new(CancellationStatus7Choice)
	return p.Cancelled
}

func (p *ProcessingStatus19Choice) AddPendingCancellation() *PendingStatus13Choice {
	p.PendingCancellation = new(PendingStatus13Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus19Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus19Choice) AddCancellationRequested() *ProprietaryReason1 {
	p.CancellationRequested = new(ProprietaryReason1)
	return p.CancellationRequested
}

func (p *ProcessingStatus19Choice) AddModificationRequested() *ProprietaryReason1 {
	p.ModificationRequested = new(ProprietaryReason1)
	return p.ModificationRequested
}
