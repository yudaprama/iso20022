package model

// Choice of formats for the specification of the status.
type Status21Choice struct {

	// Status of the transfer cancellation is accepted or sent to next party.
	Status *TransferCancellationStatus2 `xml:"Sts"`

	// Status of the transfer cancellation is rejected.
	Rejected *RejectionReason33 `xml:"Rjctd"`

	// Status of the transfer cancellation is complete. The cancellation instruction has been accepted and processed, the cancellation is complete.
	Complete *CancelledCompleteReason1 `xml:"Cmplt"`

	// Status of the transfer cancellation is pending.
	Pending *TransferCancellationPendingStatus1 `xml:"Pdg"`
}

func (s *Status21Choice) AddStatus() *TransferCancellationStatus2 {
	s.Status = new(TransferCancellationStatus2)
	return s.Status
}

func (s *Status21Choice) AddRejected() *RejectionReason33 {
	s.Rejected = new(RejectionReason33)
	return s.Rejected
}

func (s *Status21Choice) AddComplete() *CancelledCompleteReason1 {
	s.Complete = new(CancelledCompleteReason1)
	return s.Complete
}

func (s *Status21Choice) AddPending() *TransferCancellationPendingStatus1 {
	s.Pending = new(TransferCancellationPendingStatus1)
	return s.Pending
}
