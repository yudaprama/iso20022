package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus20Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus1Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus3Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus13Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus34Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus20Choice) AddCancelled() *CancelledStatus1Choice {
	i.Cancelled = new(CancelledStatus1Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus20Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus3Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus3Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus20Choice) AddRejected() *RejectedStatus13Choice {
	i.Rejected = new(RejectedStatus13Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus20Choice) AddPending() *PendingStatus34Choice {
	i.Pending = new(PendingStatus34Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus20Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus20Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus20Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
