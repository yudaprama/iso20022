package model

// Choice of format for the processing status.
type ProcessingStatus62Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus15Choice `xml:"PdgPrcg"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus16Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus46Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason7 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason5 `xml:"CxlReqd"`
}

func (p *ProcessingStatus62Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus62Choice) AddPendingProcessing() *PendingProcessingStatus15Choice {
	p.PendingProcessing = new(PendingProcessingStatus15Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus62Choice) AddRepair() *RepairStatus16Choice {
	p.Repair = new(RepairStatus16Choice)
	return p.Repair
}

func (p *ProcessingStatus62Choice) AddPendingCancellation() *PendingStatus46Choice {
	p.PendingCancellation = new(PendingStatus46Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus62Choice) AddProprietary() *ProprietaryStatusAndReason7 {
	p.Proprietary = new(ProprietaryStatusAndReason7)
	return p.Proprietary
}

func (p *ProcessingStatus62Choice) AddCancellationRequested() *ProprietaryReason5 {
	p.CancellationRequested = new(ProprietaryReason5)
	return p.CancellationRequested
}
