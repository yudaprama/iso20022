package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus14Choice struct {

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus3Choice `xml:"PdgPrcg"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason1 `xml:"CxlReqd"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus7Choice `xml:"AckdAccptd"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus7Choice `xml:"Canc"`

	// The transaction was created by the account servicer or a third party. It was not instructed directly by the account owner.
	Generated *GeneratedStatus5Choice `xml:"Gnrtd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus5Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus13Choice `xml:"PdgCxl"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason1 `xml:"ModReqd"`
}

func (i *InstructionProcessingStatus14Choice) AddPendingProcessing() *PendingProcessingStatus3Choice {
	i.PendingProcessing = new(PendingProcessingStatus3Choice)
	return i.PendingProcessing
}

func (i *InstructionProcessingStatus14Choice) AddCancellationRequested() *ProprietaryReason1 {
	i.CancellationRequested = new(ProprietaryReason1)
	return i.CancellationRequested
}

func (i *InstructionProcessingStatus14Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus7Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus7Choice)
	return i.AcknowledgedAccepted
}

func (i *InstructionProcessingStatus14Choice) AddCancelled() *CancellationStatus7Choice {
	i.Cancelled = new(CancellationStatus7Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus14Choice) AddGenerated() *GeneratedStatus5Choice {
	i.Generated = new(GeneratedStatus5Choice)
	return i.Generated
}

func (i *InstructionProcessingStatus14Choice) AddRepair() *RepairStatus5Choice {
	i.Repair = new(RepairStatus5Choice)
	return i.Repair
}

func (i *InstructionProcessingStatus14Choice) AddPendingCancellation() *PendingStatus13Choice {
	i.PendingCancellation = new(PendingStatus13Choice)
	return i.PendingCancellation
}

func (i *InstructionProcessingStatus14Choice) AddModificationRequested() *ProprietaryReason1 {
	i.ModificationRequested = new(ProprietaryReason1)
	return i.ModificationRequested
}
