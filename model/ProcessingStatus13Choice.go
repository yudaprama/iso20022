package model

// Choice of format for the processing status.
type ProcessingStatus13Choice struct {

	// The cancellation is pending processing.
	PendingCancellation *PendingStatus4Choice `xml:"PdgCxl"`

	// Instruction/Request has been rejected for further processing.
	Rejected *RejectionOrRepairStatus10Choice `xml:"Rjctd"`

	// Instruction/Request is accepted but in repair.
	Repair *RejectionOrRepairStatus10Choice `xml:"Rpr"`

	// Instruction has been acknowledged by the account servicer.
	AcknowledgedAccepted *AcknowledgedAcceptedStatus1Choice `xml:"AckdAccptd"`

	// Status that cannot be reported using one of the available standard status.
	Proprietary *ProprietaryStatusAndReason1 `xml:"Prtry"`

	// Instruction/Request will not be executed.
	Denied *DeniedStatus2Choice `xml:"Dnd"`

	// Cancellation requested executed.
	Cancelled *CancellationStatus3Choice `xml:"Canc"`
}

func (p *ProcessingStatus13Choice) AddPendingCancellation() *PendingStatus4Choice {
	p.PendingCancellation = new(PendingStatus4Choice)
	return p.PendingCancellation
}

func (p *ProcessingStatus13Choice) AddRejected() *RejectionOrRepairStatus10Choice {
	p.Rejected = new(RejectionOrRepairStatus10Choice)
	return p.Rejected
}

func (p *ProcessingStatus13Choice) AddRepair() *RejectionOrRepairStatus10Choice {
	p.Repair = new(RejectionOrRepairStatus10Choice)
	return p.Repair
}

func (p *ProcessingStatus13Choice) AddAcknowledgedAccepted() *AcknowledgedAcceptedStatus1Choice {
	p.AcknowledgedAccepted = new(AcknowledgedAcceptedStatus1Choice)
	return p.AcknowledgedAccepted
}

func (p *ProcessingStatus13Choice) AddProprietary() *ProprietaryStatusAndReason1 {
	p.Proprietary = new(ProprietaryStatusAndReason1)
	return p.Proprietary
}

func (p *ProcessingStatus13Choice) AddDenied() *DeniedStatus2Choice {
	p.Denied = new(DeniedStatus2Choice)
	return p.Denied
}

func (p *ProcessingStatus13Choice) AddCancelled() *CancellationStatus3Choice {
	p.Cancelled = new(CancellationStatus3Choice)
	return p.Cancelled
}
