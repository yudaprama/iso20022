package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus7Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus7Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *NoSpecifiedReason1 `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus14Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus3Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason1 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus7Choice) AddCancellationCompleted() *CancelledStatus7Choice {
	i.CancellationCompleted = new(CancelledStatus7Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus7Choice) AddAccepted() *NoSpecifiedReason1 {
	i.Accepted = new(NoSpecifiedReason1)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus7Choice) AddRejected() *RejectedStatus14Choice {
	i.Rejected = new(RejectedStatus14Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus7Choice) AddPendingCancellation() *PendingCancellationStatus3Choice {
	i.PendingCancellation = new(PendingCancellationStatus3Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus7Choice) AddProprietaryStatus() *ProprietaryStatusAndReason1 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason1)
	return i.ProprietaryStatus
}
