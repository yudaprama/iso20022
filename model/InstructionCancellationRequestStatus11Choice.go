package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus11Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus11Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *NoSpecifiedReason1 `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus18Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus7Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason6 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus11Choice) AddCancellationCompleted() *CancelledStatus11Choice {
	i.CancellationCompleted = new(CancelledStatus11Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus11Choice) AddAccepted() *NoSpecifiedReason1 {
	i.Accepted = new(NoSpecifiedReason1)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus11Choice) AddRejected() *RejectedStatus18Choice {
	i.Rejected = new(RejectedStatus18Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus11Choice) AddPendingCancellation() *PendingCancellationStatus7Choice {
	i.PendingCancellation = new(PendingCancellationStatus7Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus11Choice) AddProprietaryStatus() *ProprietaryStatusAndReason6 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason6)
	return i.ProprietaryStatus
}
