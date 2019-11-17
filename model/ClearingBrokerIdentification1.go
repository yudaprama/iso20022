package model

// Specifies the reference number assigned by the clearing broker. A distinction can be made between the reference for the Central Counterparty (CCP) leg and the reference for the client leg of the transaction.
type ClearingBrokerIdentification1 struct {

	// Distinguishes the client leg from the central counterparty (CCP) leg in the clearing broker identification.
	SideIndicator *SideIndicator1Code `xml:"SdInd"`

	// Specifies the identification assigned to the clearing broker.
	ClearingBrokerIdentification *Max35Text `xml:"ClrBrkrId"`
}

func (c *ClearingBrokerIdentification1) SetSideIndicator(value string) {
	c.SideIndicator = (*SideIndicator1Code)(&value)
}

func (c *ClearingBrokerIdentification1) SetClearingBrokerIdentification(value string) {
	c.ClearingBrokerIdentification = (*Max35Text)(&value)
}
