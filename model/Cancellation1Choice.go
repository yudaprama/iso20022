package model

// Choice between cancellation by reference or by transfer details.
type Cancellation1Choice struct {

	// Reference of the transfer to be cancelled.
	Reference *TransferReference1 `xml:"Ref"`

	// Details of the transfer out request to cancel.
	TransferOutDetails *TransferOut9 `xml:"TrfOutDtls"`
}

func (c *Cancellation1Choice) AddReference() *TransferReference1 {
	c.Reference = new(TransferReference1)
	return c.Reference
}

func (c *Cancellation1Choice) AddTransferOutDetails() *TransferOut9 {
	c.TransferOutDetails = new(TransferOut9)
	return c.TransferOutDetails
}
