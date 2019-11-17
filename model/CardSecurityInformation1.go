package model

// Card security code (CSC) associated with the card performing the transaction.
type CardSecurityInformation1 struct {

	// Card security code (CSC) management associated with the transaction.
	CSCManagement *CSCManagement1Code `xml:"CSCMgmt"`

	// Card security code (CSC).
	CSCValue *Min3Max4NumericText `xml:"CSCVal,omitempty"`
}

func (c *CardSecurityInformation1) SetCSCManagement(value string) {
	c.CSCManagement = (*CSCManagement1Code)(&value)
}

func (c *CardSecurityInformation1) SetCSCValue(value string) {
	c.CSCValue = (*Min3Max4NumericText)(&value)
}
