package model

// Choice between cancellation by reference or by transfer details.
type Cancellation10Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference9 `xml:"Refs"`

	// Details of the transfer in request to cancel.
	TransferInDetails *TransferIn15 `xml:"TrfInDtls"`
}

func (c *Cancellation10Choice) AddReferences() *TransferReference9 {
	newValue := new(TransferReference9)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation10Choice) AddTransferInDetails() *TransferIn15 {
	c.TransferInDetails = new(TransferIn15)
	return c.TransferInDetails
}
