package model

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason struct {

	// Status of the transfer cancellation instruction.
	Status *TransferCancellationStatus `xml:"Sts"`

	// Status of transfer cancellation is rejected.
	Rejected *TransferCancellationRejectedStatus1Choice `xml:"Rjctd"`

	// Status of the transfer cancellation is complete. The cancellation instruction has been accepted and processed, the cancellation is complete.
	Complete *TransferCancellationCompleteStatusChoice `xml:"Cmplt"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification1Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason) AddStatus() *TransferCancellationStatus {
	c.Status = new(TransferCancellationStatus)
	return c.Status
}

func (c *CancellationStatusAndReason) AddRejected() *TransferCancellationRejectedStatus1Choice {
	c.Rejected = new(TransferCancellationRejectedStatus1Choice)
	return c.Rejected
}

func (c *CancellationStatusAndReason) AddComplete() *TransferCancellationCompleteStatusChoice {
	c.Complete = new(TransferCancellationCompleteStatusChoice)
	return c.Complete
}

func (c *CancellationStatusAndReason) AddStatusInitiator() *PartyIdentification1Choice {
	c.StatusInitiator = new(PartyIdentification1Choice)
	return c.StatusInitiator
}
