package model

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason2 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer cancellation is accepted or sent to next party.
	Status *TransferCancellationStatus2 `xml:"Sts"`

	// Status of the transfer cancellation is rejected.
	Rejected *TransferCancellationRejectedStatus1 `xml:"Rjctd"`

	// Status of the transfer cancellation is complete. The cancellation instruction has been accepted and processed, the cancellation is complete.
	Complete *TransferCancellationCompleteStatusAndReason1 `xml:"Cmplt"`

	// Status of the transfer cancellation is pending.
	Pending *TransferCancellationPendingStatus1 `xml:"Pdg"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification2Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason2) SetMasterReference(value string) {
	c.MasterReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetTransferReference(value string) {
	c.TransferReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetClientReference(value string) {
	c.ClientReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) SetCancellationReference(value string) {
	c.CancellationReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason2) AddStatus() *TransferCancellationStatus2 {
	c.Status = new(TransferCancellationStatus2)
	return c.Status
}

func (c *CancellationStatusAndReason2) AddRejected() *TransferCancellationRejectedStatus1 {
	c.Rejected = new(TransferCancellationRejectedStatus1)
	return c.Rejected
}

func (c *CancellationStatusAndReason2) AddComplete() *TransferCancellationCompleteStatusAndReason1 {
	c.Complete = new(TransferCancellationCompleteStatusAndReason1)
	return c.Complete
}

func (c *CancellationStatusAndReason2) AddPending() *TransferCancellationPendingStatus1 {
	c.Pending = new(TransferCancellationPendingStatus1)
	return c.Pending
}

func (c *CancellationStatusAndReason2) AddStatusInitiator() *PartyIdentification2Choice {
	c.StatusInitiator = new(PartyIdentification2Choice)
	return c.StatusInitiator
}
