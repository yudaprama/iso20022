package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus12Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus3Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus1Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus1Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus1Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus12Choice) AddCancelled() *CancelledStatus3Choice {
	i.Cancelled = new(CancelledStatus3Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus12Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus1Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus1Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus12Choice) AddRejected() *RejectedStatus1Choice {
	i.Rejected = new(RejectedStatus1Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus12Choice) AddPending() *PendingStatus1Choice {
	i.Pending = new(PendingStatus1Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus12Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus12Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus12Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
