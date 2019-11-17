package model

// Status of the confirmation.
type ConfirmationStatus1Choice struct {

	// Status of the order confirmation is rejected.
	ConfirmationRejected []*ConfirmationRejectedStatus2 `xml:"ConfRjctd"`

	// Status of the order confirmation amendment is rejected.
	AmendmentRejected []*ConfirmationRejectedStatus2 `xml:"AmdmntRjctd"`

	// Status of the order confirmation is accepted or received or sent to next party or there is a communication problem with next party. There is no reason attached.
	Status *OrderConfirmationStatus1Code `xml:"Sts"`
}

func (c *ConfirmationStatus1Choice) AddConfirmationRejected() *ConfirmationRejectedStatus2 {
	newValue := new(ConfirmationRejectedStatus2)
	c.ConfirmationRejected = append(c.ConfirmationRejected, newValue)
	return newValue
}

func (c *ConfirmationStatus1Choice) AddAmendmentRejected() *ConfirmationRejectedStatus2 {
	newValue := new(ConfirmationRejectedStatus2)
	c.AmendmentRejected = append(c.AmendmentRejected, newValue)
	return newValue
}

func (c *ConfirmationStatus1Choice) SetStatus(value string) {
	c.Status = (*OrderConfirmationStatus1Code)(&value)
}
