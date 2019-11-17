package model

// Choice of format for the processing status.
type ProcessingStatus1Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus1Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus1Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus1Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus4Choice `xml:"Canc"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus4Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or consent.
	CancellationRequested *NoSpecifiedReason1 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *NoSpecifiedReason1 `xml:"ModReqd"`
}

func (p *ProcessingStatus1Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus1Choice) AddPendingProcessing() *PendingProcessingStatus1Choice {
	p.PendingProcessing = new(PendingProcessingStatus1Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus1Choice) AddRejected() *RejectionStatus1Choice {
	p.Rejected = new(RejectionStatus1Choice)
	return p.Rejected
}

func (p *ProcessingStatus1Choice) AddRepair() *RepairStatus1Choice {
	p.Repair = new(RepairStatus1Choice)
	return p.Repair
}

func (p *ProcessingStatus1Choice) AddCancelled() *CancellationStatus4Choice {
	p.Cancelled = new(CancellationStatus4Choice)
	return p.Cancelled
}

func (p *ProcessingStatus1Choice) AddPendingCancellation() *PendingStatus4Choice {
	p.PendingCancellation = new(PendingStatus4Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus1Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus1Choice) AddCancellationRequested() *NoSpecifiedReason1 {
	p.CancellationRequested = new(NoSpecifiedReason1)
	return p.CancellationRequested
}

func (p *ProcessingStatus1Choice) AddModificationRequested() *NoSpecifiedReason1 {
	p.ModificationRequested = new(NoSpecifiedReason1)
	return p.ModificationRequested
}
