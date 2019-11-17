package model

// Choice of format for the instruction processing status.
type InstructionProcessingStatus22Choice struct {

	// Processing of the instruction/request is pending.
	PendingProcessing *PendingProcessingStatus11Choice `xml:"PdgPrcg"`

	// Cancellation request from your counterparty for this transaction is pending waiting for your cancellation request.
	CancellationRequested *ProprietaryReason4 `xml:"CxlReqd"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus21Choice `xml:"AckdAccptd"`

	// Instruction has been cancelled.
	Cancelled *CancellationStatus14Choice `xml:"Canc"`

	// Transaction was created by the account servicer or a third party. It was not instructed directly by the account owner.
	Generated *GeneratedStatus7Choice `xml:"Gnrtd"`

	// Instruction/Request is accepted but in repair.
	Repair *RepairStatus12Choice `xml:"Rpr"`

	// Cancellation request from yourself for this instruction is pending waiting for further processing.
	PendingCancellation *PendingStatus38Choice `xml:"PdgCxl"`

	// Modification request from your counterparty for this transaction is pending waiting for your cancellation request or your consent.
	ModificationRequested *ProprietaryReason4 `xml:"ModReqd"`
}

func (i *InstructionProcessingStatus22Choice) AddPendingProcessing() *PendingProcessingStatus11Choice {
	i.PendingProcessing = new(PendingProcessingStatus11Choice)
	return i.PendingProcessing
}

func (i *InstructionProcessingStatus22Choice) AddCancellationRequested() *ProprietaryReason4 {
	i.CancellationRequested = new(ProprietaryReason4)
	return i.CancellationRequested
}

func (i *InstructionProcessingStatus22Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus21Choice {
	i.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus21Choice)
	return i.AcknowledgedAccepted
}

func (i *InstructionProcessingStatus22Choice) AddCancelled() *CancellationStatus14Choice {
	i.Cancelled = new(CancellationStatus14Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus22Choice) AddGenerated() *GeneratedStatus7Choice {
	i.Generated = new(GeneratedStatus7Choice)
	return i.Generated
}

func (i *InstructionProcessingStatus22Choice) AddRepair() *RepairStatus12Choice {
	i.Repair = new(RepairStatus12Choice)
	return i.Repair
}

func (i *InstructionProcessingStatus22Choice) AddPendingCancellation() *PendingStatus38Choice {
	i.PendingCancellation = new(PendingStatus38Choice)
	return i.PendingCancellation
}

func (i *InstructionProcessingStatus22Choice) AddModificationRequested() *ProprietaryReason4 {
	i.ModificationRequested = new(ProprietaryReason4)
	return i.ModificationRequested
}
