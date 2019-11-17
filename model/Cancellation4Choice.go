package model

// Choice between cancellation by reference or by transfer details.
type Cancellation4Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference5 `xml:"Refs"`

	// Details of the transfer out request to cancel.
	TransferOutDetails *TransferOut11 `xml:"TrfOutDtls"`
}

func (c *Cancellation4Choice) AddReferences() *TransferReference5 {
	newValue := new(TransferReference5)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation4Choice) AddTransferOutDetails() *TransferOut11 {
	c.TransferOutDetails = new(TransferOut11)
	return c.TransferOutDetails
}
