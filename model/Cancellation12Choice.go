package model

// Choice between cancellation by reference or by transfer details.
type Cancellation12Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference9 `xml:"Refs"`

	// Details of the transfer out request to cancel.
	TransferOutDetails *TransferOut17 `xml:"TrfOutDtls"`
}

func (c *Cancellation12Choice) AddReferences() *TransferReference9 {
	newValue := new(TransferReference9)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation12Choice) AddTransferOutDetails() *TransferOut17 {
	c.TransferOutDetails = new(TransferOut17)
	return c.TransferOutDetails
}
