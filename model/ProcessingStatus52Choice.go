package model

// Choice of format for the processing status.
type ProcessingStatus52Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus11Choice `xml:"PdgPrcg"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus12Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus38Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason6 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason4 `xml:"CxlReqd"`
}

func (p *ProcessingStatus52Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus52Choice) AddPendingProcessing() *PendingProcessingStatus11Choice {
	p.PendingProcessing = new(PendingProcessingStatus11Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus52Choice) AddRepair() *RepairStatus12Choice {
	p.Repair = new(RepairStatus12Choice)
	return p.Repair
}

func (p *ProcessingStatus52Choice) AddPendingCancellation() *PendingStatus38Choice {
	p.PendingCancellation = new(PendingStatus38Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus52Choice) AddProprietary() *ProprietaryStatusAndReason6 {
	p.Proprietary = new(ProprietaryStatusAndReason6)
	return p.Proprietary
}

func (p *ProcessingStatus52Choice) AddCancellationRequested() *ProprietaryReason4 {
	p.CancellationRequested = new(ProprietaryReason4)
	return p.CancellationRequested
}
