package model

// Choice between different instruction cancellation request statuses.
type InstructionCancellationRequestStatus10Choice struct {

	// Provides status information related to a instruction cancellation request completed.
	CancellationCompleted *CancelledStatus14Choice `xml:"CxlCmpltd"`

	// Provides status information related to a cancellation request accepted for further processing.
	Accepted *NoSpecifiedReason1 `xml:"Accptd"`

	// Provides status information related to a cancellation request rejected for further processing due to system (data) reasons.
	Rejected *RejectedStatus20Choice `xml:"Rjctd"`

	// Provides status information related to a pending cancellation request.
	PendingCancellation *PendingCancellationStatus6Choice `xml:"PdgCxl"`

	// Proprietary status related to an instruction cancellation request.
	ProprietaryStatus *ProprietaryStatusAndReason7 `xml:"PrtrySts"`
}

func (i *InstructionCancellationRequestStatus10Choice) AddCancellationCompleted() *CancelledStatus14Choice {
	i.CancellationCompleted = new(CancelledStatus14Choice)
	return i.CancellationCompleted
}

func (i *InstructionCancellationRequestStatus10Choice) AddAccepted() *NoSpecifiedReason1 {
	i.Accepted = new(NoSpecifiedReason1)
	return i.Accepted
}

func (i *InstructionCancellationRequestStatus10Choice) AddRejected() *RejectedStatus20Choice {
	i.Rejected = new(RejectedStatus20Choice)
	return i.Rejected
}

func (i *InstructionCancellationRequestStatus10Choice) AddPendingCancellation() *PendingCancellationStatus6Choice {
	i.PendingCancellation = new(PendingCancellationStatus6Choice)
	return i.PendingCancellation
}

func (i *InstructionCancellationRequestStatus10Choice) AddProprietaryStatus() *ProprietaryStatusAndReason7 {
	i.ProprietaryStatus = new(ProprietaryStatusAndReason7)
	return i.ProprietaryStatus
}
