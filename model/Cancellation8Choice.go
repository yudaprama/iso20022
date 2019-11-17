package model

// Choice between cancellation by reference or by transfer details.
type Cancellation8Choice struct {

	// Reference of the transfer to be cancelled.
	References []*TransferReference5 `xml:"Refs"`

	// Details of the transfer out request to cancel.
	TransferOutDetails *TransferOut15 `xml:"TrfOutDtls"`
}

func (c *Cancellation8Choice) AddReferences() *TransferReference5 {
	newValue := new(TransferReference5)
	c.References = append(c.References, newValue)
	return newValue
}

func (c *Cancellation8Choice) AddTransferOutDetails() *TransferOut15 {
	c.TransferOutDetails = new(TransferOut15)
	return c.TransferOutDetails
}
