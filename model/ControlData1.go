package model

// Provides details of the number of transactions and the control sum of the message.
type ControlData1 struct {

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`
}

func (c *ControlData1) SetNumberOfTransactions(value string) {
	c.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (c *ControlData1) SetControlSum(value string) {
	c.ControlSum = (*DecimalNumber)(&value)
}
