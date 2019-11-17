package model

// Choice between cancellation by reference or by transfer details.
type Cancellation5Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference5 `xml:"Refs"`

	// Details of the transfer in request to cancel.
	TransferInDetails *TransferIn10 `xml:"TrfInDtls"`
}

func (c *Cancellation5Choice) AddReferences() *TransferReference5 {
	newValue := new(TransferReference5)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation5Choice) AddTransferInDetails() *TransferIn10 {
	c.TransferInDetails = new(TransferIn10)
	return c.TransferInDetails
}
