package model

// Choice between cancellation by transfer details or reference.
type Cancellation6Choice struct {

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *ISATransfer12 `xml:"CxlByTrfInstrDtls"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *TransferReference7 `xml:"CxlByRef"`
}

func (c *Cancellation6Choice) AddCancellationByTransferInstructionDetails() *ISATransfer12 {
	c.CancellationByTransferInstructionDetails = new(ISATransfer12)
	return c.CancellationByTransferInstructionDetails
}

func (c *Cancellation6Choice) AddCancellationByReference() *TransferReference7 {
	c.CancellationByReference = new(TransferReference7)
	return c.CancellationByReference
}
