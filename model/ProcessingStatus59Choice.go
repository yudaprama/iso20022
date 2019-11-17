package model

// Provides the status of a transaction.
type ProcessingStatus59Choice struct {

	// Instruction is rejected.
	Rejected *RejectedStatus22Choice `xml:"Rjctd"`

	// Instruction is cancelled.
	Cancelled *CancelledStatus16Choice `xml:"Canc"`

	// Instruction is accepted.
	Accepted *AcceptedStatus10Choice `xml:"Accptd"`
}

func (p *ProcessingStatus59Choice) AddRejected() *RejectedStatus22Choice {
	p.Rejected = new(RejectedStatus22Choice)
	return p.Rejected
}

func (p *ProcessingStatus59Choice) AddCancelled() *CancelledStatus16Choice {
	p.Cancelled = new(CancelledStatus16Choice)
	return p.Cancelled
}

func (p *ProcessingStatus59Choice) AddAccepted() *AcceptedStatus10Choice {
	p.Accepted = new(AcceptedStatus10Choice)
	return p.Accepted
}
