package model

// Choice of format for the processing status.
type ProcessingStatus23Choice struct {

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus3Choice `xml:"PdgPrcg"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus5Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus13Choice `xml:"PdgCxl"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason1 `xml:"CxlReqd"`
}

func (p *ProcessingStatus23Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus23Choice) AddPendingProcessing() *PendingProcessingStatus3Choice {
	p.PendingProcessing = new(PendingProcessingStatus3Choice)
	return p.PendingProcessing
}

func (p *ProcessingStatus23Choice) AddRepair() *RepairStatus5Choice {
	p.Repair = new(RepairStatus5Choice)
	return p.Repair
}

func (p *ProcessingStatus23Choice) AddPendingCancellation() *PendingStatus13Choice {
	p.PendingCancellation = new(PendingStatus13Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus23Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus23Choice) AddCancellationRequested() *ProprietaryReason1 {
	p.CancellationRequested = new(ProprietaryReason1)
	return p.CancellationRequested
}
