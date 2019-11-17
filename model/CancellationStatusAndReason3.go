package model

// Status of a transfer cancellation instruction and the reason for the status.
type CancellationStatusAndReason3 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identification of a transfer, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *AdditionalReference7 `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer cancellation, as assigned by the instructing party.
	CancellationReference *Max35Text `xml:"CxlRef,omitempty"`

	// Status of the transfer cancellation.
	Status *Status21Choice `xml:"Sts"`

	// Party that initiates the status.
	StatusInitiator *PartyIdentification70Choice `xml:"StsInitr,omitempty"`
}

func (c *CancellationStatusAndReason3) SetMasterReference(value string) {
	c.MasterReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) SetTransferReference(value string) {
	c.TransferReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) AddClientReference() *AdditionalReference7 {
	c.ClientReference = new(AdditionalReference7)
	return c.ClientReference
}

func (c *CancellationStatusAndReason3) SetCancellationReference(value string) {
	c.CancellationReference = (*Max35Text)(&value)
}

func (c *CancellationStatusAndReason3) AddStatus() *Status21Choice {
	c.Status = new(Status21Choice)
	return c.Status
}

func (c *CancellationStatusAndReason3) AddStatusInitiator() *PartyIdentification70Choice {
	c.StatusInitiator = new(PartyIdentification70Choice)
	return c.StatusInitiator
}
