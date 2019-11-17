package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus27Choice struct {

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus15Choice `xml:"PdgPrcg"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason5 `xml:"CxlReqd"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus25Choice `xml:"AckdAccptd"`

	// Instruction has been cancelled.
	Cancelled *CancellationStatus17Choice `xml:"Canc"`

	// Transaction was created by the account servicer or a third party. It was not instructed directly by the account owner.
	Generated *GeneratedStatus8Choice `xml:"Gnrtd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus16Choice `xml:"Rpr"`

	// Cancellation request from yourself for this instruction is pending waiting for further processing.
	PendingCancellation *PendingStatus46Choice `xml:"PdgCxl"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason5 `xml:"ModReqd"`
}

func (i *InstructionProcessingStatus27Choice) AddPendingProcessing() *PendingProcessingStatus15Choice {
	i.PendingProcessing = new(PendingProcessingStatus15Choice)
	return i.PendingProcessing
}

func (i *InstructionProcessingStatus27Choice) AddCancellationRequested() *ProprietaryReason5 {
	i.CancellationRequested = new(ProprietaryReason5)
	return i.CancellationRequested
}

func (i *InstructionProcessingStatus27Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus25Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus25Choice)
	return i.AcknowledgedAccepted
}

func (i *InstructionProcessingStatus27Choice) AddCancelled() *CancellationStatus17Choice {
	i.Cancelled = new(CancellationStatus17Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus27Choice) AddGenerated() *GeneratedStatus8Choice {
	i.Generated = new(GeneratedStatus8Choice)
	return i.Generated
}

func (i *InstructionProcessingStatus27Choice) AddRepair() *RepairStatus16Choice {
	i.Repair = new(RepairStatus16Choice)
	return i.Repair
}

func (i *InstructionProcessingStatus27Choice) AddPendingCancellation() *PendingStatus46Choice {
	i.PendingCancellation = new(PendingStatus46Choice)
	return i.PendingCancellation
}

func (i *InstructionProcessingStatus27Choice) AddModificationRequested() *ProprietaryReason5 {
	i.ModificationRequested = new(ProprietaryReason5)
	return i.ModificationRequested
}
