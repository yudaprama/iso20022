package model

// Specifies the corrective transaction on which the investigation is processed.
type CorrectiveTransaction2Choice struct {

	// Set of elements used to reference the details of the corrective payment initiation.
	Initiation *CorrectivePaymentInitiation2 `xml:"Initn"`

	// Set of elements used to reference the details of the corrective interbank payment transaction.
	Interbank *CorrectiveInterbankTransaction1 `xml:"IntrBk"`
}

func (c *CorrectiveTransaction2Choice) AddInitiation() *CorrectivePaymentInitiation2 {
	c.Initiation = new(CorrectivePaymentInitiation2)
	return c.Initiation
}

func (c *CorrectiveTransaction2Choice) AddInterbank() *CorrectiveInterbankTransaction1 {
	c.Interbank = new(CorrectiveInterbankTransaction1)
	return c.Interbank
}
