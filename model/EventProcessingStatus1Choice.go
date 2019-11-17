package model

// Provides information about the status of a corporate action or the status of a payment.
type EventProcessingStatus1Choice struct {

	// Specifies that a corporate action event processing has been completed.
	Complete *NoSpecifiedReason1 `xml:"Cmplt"`

	// Corporate action event processing specifying that the funds paid have been reconciled with the funds received from the agent (meaning that there is no more risk of payment to be reversed).
	Reconciled *NoSpecifiedReason1 `xml:"Rcncld"`

	// Specifies that a corporate action event processing has not been fully completed and is therefore pending.
	Pending *PendingStatus2Choice `xml:"Pdg"`

	// Proprietary status related to the event processing.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (e *EventProcessingStatus1Choice) AddComplete() *NoSpecifiedReason1 {
	e.Complete = new(NoSpecifiedReason1)
	return e.Complete
}

func (e *EventProcessingStatus1Choice) AddReconciled() *NoSpecifiedReason1 {
	e.Reconciled = new(NoSpecifiedReason1)
	return e.Reconciled
}

func (e *EventProcessingStatus1Choice) AddPending() *PendingStatus2Choice {
	e.Pending = new(PendingStatus2Choice)
	return e.Pending
}

func (e *EventProcessingStatus1Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	e.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return e.ProprietaryStatus
}
