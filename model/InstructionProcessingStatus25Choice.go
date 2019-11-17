package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus25Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus15Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus9Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus21Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus44Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason7 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus25Choice) AddCancelled() *CancelledStatus15Choice {
	i.Cancelled = new(CancelledStatus15Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus25Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus9Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus9Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus25Choice) AddRejected() *RejectedStatus21Choice {
	i.Rejected = new(RejectedStatus21Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus25Choice) AddPending() *PendingStatus44Choice {
	i.Pending = new(PendingStatus44Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus25Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus25Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus25Choice) AddProprietaryStatus() *ProprietaryStatusAndReason7 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason7)
	return i.ProprietaryStatus
}
