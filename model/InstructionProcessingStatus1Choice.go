package model

// Choice between different instruction processing statuses.
type InstructionProcessingStatus1Choice struct {

	// Provides status information related to an instruction request cancelled.
	Cancelled *CancelledStatus1Choice `xml:"Canc"`

	// Provides status information related to an instruction request accepted for further processing.
	Accepted *AcceptedStatus1Choice `xml:"Accptd"`

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

func (i *InstructionProcessingStatus1Choice) AddCancelled() *CancelledStatus1Choice {
	i.Cancelled = new(CancelledStatus1Choice)
	return i.Cancelled
}

func (i *InstructionProcessingStatus1Choice) AddAccepted() *AcceptedStatus1Choice {
	i.Accepted = new(AcceptedStatus1Choice)
	return i.Accepted
}

func (i *InstructionProcessingStatus1Choice) AddRejected() *RejectedStatus1Choice {
	i.Rejected = new(RejectedStatus1Choice)
	return i.Rejected
}

func (i *InstructionProcessingStatus1Choice) AddPending() *PendingStatus1Choice {
	i.Pending = new(PendingStatus1Choice)
	return i.Pending
}

func (i *InstructionProcessingStatus1Choice) AddDefaultAction() *NoSpecifiedReason1 {
	i.DefaultAction = new(NoSpecifiedReason1)
	return i.DefaultAction
}

func (i *InstructionProcessingStatus1Choice) AddStandingInstruction() *NoSpecifiedReason1 {
	i.StandingInstruction = new(NoSpecifiedReason1)
	return i.StandingInstruction
}

func (i *InstructionProcessingStatus1Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
