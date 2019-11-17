package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus29Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus12Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus8Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus23Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus52Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason6 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus29Choice) AddCancelled() *CancelledStatus12Choice {
	i.Cancelled = new(CancelledStatus12Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus29Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus8Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus8Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus29Choice) AddRejected() *RejectedStatus23Choice {
	i.Rejected = new(RejectedStatus23Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus29Choice) AddPending() *PendingStatus52Choice {
	i.Pending = new(PendingStatus52Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus29Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus29Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus29Choice) AddProprietaryStatus() *ProprietaryStatusAndReason6 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason6)
	return i.ProprietaryStatus
}
