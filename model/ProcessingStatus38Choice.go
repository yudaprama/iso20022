package model

// Choice of format for the processing status.
type ProcessingStatus38Choice struct {

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus13Choice `xml:"PdgCxl"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus4Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus14Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus6Choice `xml:"Rpr"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus8Choice `xml:"Canc"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	CancellationRequested *ProprietaryReason1 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason1 `xml:"ModReqd"`
}

func (p *ProcessingStatus38Choice) AddPendingCancellation() *PendingStatus13Choice {
	p.PendingCancellation = new(PendingStatus13Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus38Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus38Choice) AddPendingProcessing() *PendingProcessingStatus4Choice {
	p.PendingProcessing = new(PendingProcessingStatus4Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus38Choice) AddRejected() *RejectionStatus14Choice {
	p.Rejected = new(RejectionStatus14Choice)
	return p.Rejected
}

func (p *ProcessingStatus38Choice) AddRepair() *RepairStatus6Choice {
	p.Repair = new(RepairStatus6Choice)
	return p.Repair
}

func (p *ProcessingStatus38Choice) AddCancelled() *CancellationStatus8Choice {
	p.Cancelled = new(CancellationStatus8Choice)
	return p.Cancelled
}

func (p *ProcessingStatus38Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus38Choice) AddCancellationRequested() *ProprietaryReason1 {
	p.CancellationRequested = new(ProprietaryReason1)
	return p.CancellationRequested
}

func (p *ProcessingStatus38Choice) AddModificationRequested() *ProprietaryReason1 {
	p.ModificationRequested = new(ProprietaryReason1)
	return p.ModificationRequested
}
