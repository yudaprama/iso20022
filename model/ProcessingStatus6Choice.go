package model

// Choice of format for the processing status.
type ProcessingStatus6Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus1Choice `xml:"PdgPrcg"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus1Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus4Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *NoSpecifiedReason1 `xml:"CxlReqd"`
}

func (p *ProcessingStatus6Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus6Choice) AddPendingProcessing() *PendingProcessingStatus1Choice {
	p.PendingProcessing = new(PendingProcessingStatus1Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus6Choice) AddRepair() *RepairStatus1Choice {
	p.Repair = new(RepairStatus1Choice)
	return p.Repair
}

func (p *ProcessingStatus6Choice) AddPendingCancellation() *PendingStatus4Choice {
	p.PendingCancellation = new(PendingStatus4Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus6Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus6Choice) AddCancellationRequested() *NoSpecifiedReason1 {
	p.CancellationRequested = new(NoSpecifiedReason1)
	return p.CancellationRequested
}
