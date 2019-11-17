package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus3Choice struct {

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus1Choice `xml:"PdgPrcg"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *NoSpecifiedReason1 `xml:"CxlReqd"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus3Choice `xml:"AckdAccptd"`

	// Instruction has been cancelled (only as an response to an SecuritiesTransactionStatusQuery). The status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	Cancelled *CancellationStatus4Choice `xml:"Canc"`

	// The transaction was created by the account servicer or a third party. It was not instructed directly by the account owner.
	Generated *GeneratedStatus1Choice `xml:"Gnrtd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus1Choice `xml:"Rpr"`

	// A cancellation request from yourself for this instruction is pending waiting for further processing (only as an response to an SecuritiesTransactionStatusQuery). The pending status on the processing of a cancellation request must be provided using a SecuritiesTransactionCancellationRequestStatusAdvice.
	PendingCancellation *PendingStatus4Choice `xml:"PdgCxl"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *NoSpecifiedReason1 `xml:"ModReqd"`
}

func (i *InstructionProcessingStatus3Choice) AddPendingProcessing() *PendingProcessingStatus1Choice {
	i.PendingProcessing = new(PendingProcessingStatus1Choice)
	return i.PendingProcessing
}

func (i *InstructionProcessingStatus3Choice) AddCancellationRequested() *NoSpecifiedReason1 {
	i.CancellationRequested = new(NoSpecifiedReason1)
	return i.CancellationRequested
}

func (i *InstructionProcessingStatus3Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus3Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus3Choice)
	return i.AcknowledgedAccepted
}

func (i *InstructionProcessingStatus3Choice) AddCancelled() *CancellationStatus4Choice {
	i.Cancelled = new(CancellationStatus4Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus3Choice) AddGenerated() *GeneratedStatus1Choice {
	i.Generated = new(GeneratedStatus1Choice)
	return i.Generated
}

func (i *InstructionProcessingStatus3Choice) AddRepair() *RepairStatus1Choice {
	i.Repair = new(RepairStatus1Choice)
	return i.Repair
}

func (i *InstructionProcessingStatus3Choice) AddPendingCancellation() *PendingStatus4Choice {
	i.PendingCancellation = new(PendingStatus4Choice)
	return i.PendingCancellation
}

func (i *InstructionProcessingStatus3Choice) AddModificationRequested() *NoSpecifiedReason1 {
	i.ModificationRequested = new(NoSpecifiedReason1)
	return i.ModificationRequested
}
