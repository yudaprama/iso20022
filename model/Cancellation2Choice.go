package model

// Choice between cancellation by reference or by transfer details.
type Cancellation2Choice struct {

	// Reference of the transfer to be cancelled.
	Reference *TransferReference1 `xml:"Ref"`

	// Details of the transfer in request to cancel.
	TransferInDetails *TransferIn7 `xml:"TrfInDtls"`
}

func (c *Cancellation2Choice) AddReference() *TransferReference1 {
	c.Reference = new(TransferReference1)
	return c.Reference
}

func (c *Cancellation2Choice) AddTransferInDetails() *TransferIn7 {
	c.TransferInDetails = new(TransferIn7)
	return c.TransferInDetails
}
