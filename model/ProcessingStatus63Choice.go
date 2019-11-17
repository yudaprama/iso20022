package model

// Choice of format for the processing status.
type ProcessingStatus63Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus15Choice `xml:"PdgPrcg"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionStatus26Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus16Choice `xml:"Rpr"`

	// Instruction/Request has been cancelled.
	Cancelled *CancellationStatus17Choice `xml:"Canc"`

	// Cancellation request from yourself for this instruction is pending, waiting for further processing.
	PendingCancellation *PendingStatus46Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request or consent.
	CancellationRequested *ProprietaryReason5 `xml:"CxlReqd"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason5 `xml:"ModReqd"`
}

func (p *ProcessingStatus63Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus63Choice) AddPendingProcessing() *PendingProcessingStatus15Choice {
	p.PendingProcessing = new(PendingProcessingStatus15Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus63Choice) AddRejected() *RejectionStatus26Choice {
	p.Rejected = new(RejectionStatus26Choice)
	return p.Rejected
}

func (p *ProcessingStatus63Choice) AddRepair() *RepairStatus16Choice {
	p.Repair = new(RepairStatus16Choice)
	return p.Repair
}

func (p *ProcessingStatus63Choice) AddCancelled() *CancellationStatus17Choice {
	p.Cancelled = new(CancellationStatus17Choice)
	return p.Cancelled
}

func (p *ProcessingStatus63Choice) AddPendingCancellation() *PendingStatus46Choice {
	p.PendingCancellation = new(PendingStatus46Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus63Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}

func (p *ProcessingStatus63Choice) AddCancellationRequested() *ProprietaryReason5 {
	p.CancellationRequested = new(ProprietaryReason5)
	return p.CancellationRequested
}

func (p *ProcessingStatus63Choice) AddModificationRequested() *ProprietaryReason5 {
	p.ModificationRequested = new(ProprietaryReason5)
	return p.ModificationRequested
}
