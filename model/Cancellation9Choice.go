package model

// Choice between cancellation by reference or by transfer details.
type Cancellation9Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference5 `xml:"Refs"`

	// Details of the transfer in request to cancel.
	TransferInDetails *TransferIn13 `xml:"TrfInDtls"`
}

func (c *Cancellation9Choice) AddReferences() *TransferReference5 {
	newValue := new(TransferReference5)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation9Choice) AddTransferInDetails() *TransferIn13 {
	c.TransferInDetails = new(TransferIn13)
	return c.TransferInDetails
}
