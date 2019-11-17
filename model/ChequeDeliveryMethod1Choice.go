package model

// Set of characteristics related to a cheque instruction, such as cheque type or cheque number.
type ChequeDeliveryMethod1Choice struct {

	// Specifies the delivery method of the cheque by the debtor's agent.
	Code *ChequeDelivery1Code `xml:"Cd"`

	// Specifies a proprietary delivery method of the cheque by the debtor's agent.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *ChequeDeliveryMethod1Choice) SetCode(value string) {
	c.Code = (*ChequeDelivery1Code)(&value)
}

func (c *ChequeDeliveryMethod1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
