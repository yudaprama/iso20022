package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus24Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus12Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus8Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus19Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus42Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason6 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus24Choice) AddCancelled() *CancelledStatus12Choice {
	i.Cancelled = new(CancelledStatus12Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus24Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus8Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus8Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus24Choice) AddRejected() *RejectedStatus19Choice {
	i.Rejected = new(RejectedStatus19Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus24Choice) AddPending() *PendingStatus42Choice {
	i.Pending = new(PendingStatus42Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus24Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus24Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus24Choice) AddProprietaryStatus() *ProprietaryStatusAndReason6 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason6)
	return i.ProprietaryStatus
}
