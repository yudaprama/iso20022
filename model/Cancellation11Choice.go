package model

// Choice between cancellation by transfer details or reference.
type Cancellation11Choice struct {

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *ISATransfer24 `xml:"CxlByTrfInstrDtls"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *TransferReference7 `xml:"CxlByRef"`
}

func (c *Cancellation11Choice) AddCancellationByTransferInstructionDetails() *ISATransfer24 {
	c.CancellationByTransferInstructionDetails = new(ISATransfer24)
	return c.CancellationByTransferInstructionDetails
}

func (c *Cancellation11Choice) AddCancellationByReference() *TransferReference7 {
	c.CancellationByReference = new(TransferReference7)
	return c.CancellationByReference
}
