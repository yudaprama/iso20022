package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus30Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus15Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus9Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus24Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus53Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason7 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus30Choice) AddCancelled() *CancelledStatus15Choice {
	i.Cancelled = new(CancelledStatus15Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus30Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus9Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus9Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus30Choice) AddRejected() *RejectedStatus24Choice {
	i.Rejected = new(RejectedStatus24Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus30Choice) AddPending() *PendingStatus53Choice {
	i.Pending = new(PendingStatus53Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus30Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus30Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus30Choice) AddProprietaryStatus() *ProprietaryStatusAndReason7 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason7)
	return i.ProprietaryStatus
}
