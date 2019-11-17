package model

// Choice between cancellation by transfer details or reference.
type Cancellation7Choice struct {

	// Information related to the transfer instruction to be cancelled.
	CancellationByTransferInstructionDetails *ISATransfer19 `xml:"CxlByTrfInstrDtls"`

	// Reference of the transfer instruction to be cancelled.
	CancellationByReference *TransferReference7 `xml:"CxlByRef"`
}

func (c *Cancellation7Choice) AddCancellationByTransferInstructionDetails() *ISATransfer19 {
	c.CancellationByTransferInstructionDetails = new(ISATransfer19)
	return c.CancellationByTransferInstructionDetails
}

func (c *Cancellation7Choice) AddCancellationByReference() *TransferReference7 {
	c.CancellationByReference = new(TransferReference7)
	return c.CancellationByReference
}
