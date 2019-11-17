package model

// Provides the status of a transaction (eg, at a non-matching CSD) as far as the message sender is concerned.
type ProcessingStatus39Choice struct {

	// Instruction is rejected.
	Rejected *RejectedStatus10Choice `xml:"Rjctd"`

	// Instruction is cancelled.
	Cancelled *CancelledStatus5Choice `xml:"Canc"`

	// Instruction is accepted.
	Accepted *AcceptedStatus4Choice `xml:"Accptd"`
}

func (p *ProcessingStatus39Choice) AddRejected() *RejectedStatus10Choice {
	p.Rejected = new(RejectedStatus10Choice)
	return p.Rejected
}

func (p *ProcessingStatus39Choice) AddCancelled() *CancelledStatus5Choice {
	p.Cancelled = new(CancelledStatus5Choice)
	return p.Cancelled
}

func (p *ProcessingStatus39Choice) AddAccepted() *AcceptedStatus4Choice {
	p.Accepted = new(AcceptedStatus4Choice)
	return p.Accepted
}
