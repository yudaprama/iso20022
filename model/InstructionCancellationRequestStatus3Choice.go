package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus3Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus3Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *AcceptedStatus1Choice `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus1Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus1Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus3Choice) AddCancellationCompleted() *CancelledStatus3Choice {
	i.CancellationCompleted = new(CancelledStatus3Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus3Choice) AddAccepted() *AcceptedStatus1Choice {
	i.Accepted = new(AcceptedStatus1Choice)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus3Choice) AddRejected() *RejectedStatus1Choice {
	i.Rejected = new(RejectedStatus1Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus3Choice) AddPendingCancellation() *PendingCancellationStatus1Choice {
	i.PendingCancellation = new(PendingCancellationStatus1Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus3Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
