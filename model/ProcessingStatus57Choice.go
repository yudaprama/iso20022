package model

// Choice of format for the processing status.
type ProcessingStatus57Choice struct {

	// Cancellation request from yourself for this instruction is pending waiting for further processing.
	PendingCancellation *PendingStatus46Choice `xml:"PdgCxl"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus14Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus22Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus15Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus18Choice `xml:"Canc"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	CancellationRequested *ProprietaryReason5 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason5 `xml:"ModReqd"`
}

func (p *ProcessingStatus57Choice) AddPendingCancellation() *PendingStatus46Choice {
	p.PendingCancellation = new(PendingStatus46Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus57Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus57Choice) AddPendingProcessing() *PendingProcessingStatus14Choice {
	p.PendingProcessing = new(PendingProcessingStatus14Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus57Choice) AddRejected() *RejectionStatus22Choice {
	p.Rejected = new(RejectionStatus22Choice)
	return p.Rejected
}

func (p *ProcessingStatus57Choice) AddRepair() *RepairStatus15Choice {
	p.Repair = new(RepairStatus15Choice)
	return p.Repair
}

func (p *ProcessingStatus57Choice) AddCancelled() *CancellationStatus18Choice {
	p.Cancelled = new(CancellationStatus18Choice)
	return p.Cancelled
}

func (p *ProcessingStatus57Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}

func (p *ProcessingStatus57Choice) AddCancellationRequested() *ProprietaryReason5 {
	p.CancellationRequested = new(ProprietaryReason5)
	return p.CancellationRequested
}

func (p *ProcessingStatus57Choice) AddModificationRequested() *ProprietaryReason5 {
	p.ModificationRequested = new(ProprietaryReason5)
	return p.ModificationRequested
}
