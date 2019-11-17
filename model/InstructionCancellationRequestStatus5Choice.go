package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus5Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus3Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *AcceptedStatus1Choice `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus1Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus3Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus5Choice) AddCancellationCompleted() *CancelledStatus3Choice {
	i.CancellationCompleted = new(CancelledStatus3Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus5Choice) AddAccepted() *AcceptedStatus1Choice {
	i.Accepted = new(AcceptedStatus1Choice)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus5Choice) AddRejected() *RejectedStatus1Choice {
	i.Rejected = new(RejectedStatus1Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus5Choice) AddPendingCancellation() *PendingCancellationStatus3Choice {
	i.PendingCancellation = new(PendingCancellationStatus3Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus5Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
