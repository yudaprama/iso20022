package model

// Specifies the corrective transaction on which the investigation is processed.
type CorrectiveTransaction1Choice struct {

	// Set of elements used to reference the details of the corrective payment initiation.
	Initiation *CorrectivePaymentInitiation1 `xml:"Initn"`

	// Set of elements used to reference the details of the corrective interbank payment transaction.
	Interbank *CorrectiveInterbankTransaction1 `xml:"IntrBk"`
}

func (c *CorrectiveTransaction1Choice) AddInitiation() *CorrectivePaymentInitiation1 {
	c.Initiation = new(CorrectivePaymentInitiation1)
	return c.Initiation
}

func (c *CorrectiveTransaction1Choice) AddInterbank() *CorrectiveInterbankTransaction1 {
	c.Interbank = new(CorrectiveInterbankTransaction1)
	return c.Interbank
}
