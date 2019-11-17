package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus9Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus11Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *NoSpecifiedReason1 `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus18Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus5Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason6 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus9Choice) AddCancellationCompleted() *CancelledStatus11Choice {
	i.CancellationCompleted = new(CancelledStatus11Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus9Choice) AddAccepted() *NoSpecifiedReason1 {
	i.Accepted = new(NoSpecifiedReason1)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus9Choice) AddRejected() *RejectedStatus18Choice {
	i.Rejected = new(RejectedStatus18Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus9Choice) AddPendingCancellation() *PendingCancellationStatus5Choice {
	i.PendingCancellation = new(PendingCancellationStatus5Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus9Choice) AddProprietaryStatus() *ProprietaryStatusAndReason6 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason6)
	return i.ProprietaryStatus
}
