package model

// Choice between cancellation by transfer details or reference.
type Cancellation3Choice struct {

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *ISATransfer9 `xml:"CxlByTrfInstrDtls"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *TransferReference3 `xml:"CxlByRef"`
}

func (c *Cancellation3Choice) AddCancellationByTransferInstructionDetails() *ISATransfer9 {
	c.CancellationByTransferInstructionDetails = new(ISATransfer9)
	return c.CancellationByTransferInstructionDetails
}

func (c *Cancellation3Choice) AddCancellationByReference() *TransferReference3 {
	c.CancellationByReference = new(TransferReference3)
	return c.CancellationByReference
}
