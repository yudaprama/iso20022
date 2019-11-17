package model

// Provides the status of a transaction.
type ProcessingStatus56Choice struct {

	// Instruction is rejected.
	Rejected *RejectedStatus17Choice `xml:"Rjctd"`

	// Instruction is cancelled.
	Cancelled *CancelledStatus10Choice `xml:"Canc"`

	// Instruction is accepted.
	Accepted *AcceptedStatus7Choice `xml:"Accptd"`
}

func (p *ProcessingStatus56Choice) AddRejected() *RejectedStatus17Choice {
	p.Rejected = new(RejectedStatus17Choice)
	return p.Rejected
}

func (p *ProcessingStatus56Choice) AddCancelled() *CancelledStatus10Choice {
	p.Cancelled = new(CancelledStatus10Choice)
	return p.Cancelled
}

func (p *ProcessingStatus56Choice) AddAccepted() *AcceptedStatus7Choice {
	p.Accepted = new(AcceptedStatus7Choice)
	return p.Accepted
}
