package model

// Provides information about the status of a corporate action or the status of a payment.
type EventProcessingStatus3Choice struct {

	// Specifies that a corporate action event processing has been completed.
	Complete *NoSpecifiedReason1 `xml:"Cmplt"`

	// Corporate action event processing specifying that the funds paid have been reconciled with the funds received from the agent (meaning that there is no more risk of payment to be reversed).
	Reconciled *NoSpecifiedReason1 `xml:"Rcncld"`

	// Specifies that a corporate action event processing has not been fully completed and is therefore pending.
	Pending *PendingStatus41Choice `xml:"Pdg"`

	// Proprietary status related to the event processing.
	ProprietaryStatus *ProprietaryStatusAndReason6 `xml:"PrtrySts"`
}

func (e *EventProcessingStatus3Choice) AddComplete() *NoSpecifiedReason1 {
	e.Complete = new(NoSpecifiedReason1)
	return e.Complete
}

func (e *EventProcessingStatus3Choice) AddReconciled() *NoSpecifiedReason1 {
	e.Reconciled = new(NoSpecifiedReason1)
	return e.Reconciled
}

func (e *EventProcessingStatus3Choice) AddPending() *PendingStatus41Choice {
	e.Pending = new(PendingStatus41Choice)
	return e.Pending
}

func (e *EventProcessingStatus3Choice) AddProprietaryStatus() *ProprietaryStatusAndReason6 {
	e.ProprietaryStatus = new(ProprietaryStatusAndReason6)
	return e.ProprietaryStatus
}
