package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus1Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus1Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *AcceptedStatus1Choice `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus1Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus1Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus1Choice) AddCancellationCompleted() *CancelledStatus1Choice {
	i.CancellationCompleted = new(CancelledStatus1Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus1Choice) AddAccepted() *AcceptedStatus1Choice {
	i.Accepted = new(AcceptedStatus1Choice)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus1Choice) AddRejected() *RejectedStatus1Choice {
	i.Rejected = new(RejectedStatus1Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus1Choice) AddPendingCancellation() *PendingCancellationStatus1Choice {
	i.PendingCancellation = new(PendingCancellationStatus1Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus1Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
