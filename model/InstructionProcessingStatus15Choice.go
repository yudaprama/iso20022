package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus15Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus3Choice `xml:"Canc"`

	// Provides status information related to an instruction request that is accepted. This means that the instruction has been received, is processable and has been validated for further processing.
	AcceptedForFurtherProcessing *AcceptedStatus3Choice `xml:"AccptdForFrthrPrcg"`

	// Provides status information related to an instruction request rejected for further processing due to system reasons.
	Rejected *RejectedStatus9Choice `xml:"Rjctd"`

	// Provides status information related to a pending instruction.
	Pending *PendingStatus32Choice `xml:"Pdg"`

	// Default action is taken.
	DefaultAction *NoSpecifiedReason1 `xml:"DfltActn"`

	// Standing instruction has been applied.
	StandingInstruction *NoSpecifiedReason1 `xml:"StgInstr"`

	// Proprietary status related to an instruction.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionProcessingStatus15Choice) AddCancelled() *CancelledStatus3Choice {
	i.Cancelled = new(CancelledStatus3Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus15Choice) AddAcceptedForFurtherProcessing() *AcceptedStatus3Choice {
	i.AcceptedForFurtherProcessing = new(AcceptedStatus3Choice)
	return i.AcceptedForFurtherProcessing
}

func (i *InstructionProcessingStatus15Choice) AddRejected() *RejectedStatus9Choice {
	i.Rejected = new(RejectedStatus9Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus15Choice) AddPending() *PendingStatus32Choice {
	i.Pending = new(PendingStatus32Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus15Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus15Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus15Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
