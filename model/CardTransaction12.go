package model

// Network management transaction.
type CardTransaction12 struct {

	// Type of network management service (correspond to the ISO 8583 field 24)
	NetworkManagementType *CardServiceType2Code `xml:"NtwkMgmtTp"`

	// Date and time of the transaction.
	InitiatorDateTime *ISODateTime `xml:"InitrDtTm,omitempty"`

	// Number of messages in the store and forward queue.
	NumberOfMessages *Number `xml:"NbOfMsgs,omitempty"`

	// Maximum number of messages in the store and forward queue.
	MaximumNumberOfMessages *Number `xml:"MaxNbOfMsgs,omitempty"`

	// Response to the network management request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`
}

func (c *CardTransaction12) SetNetworkManagementType(value string) {
	c.NetworkManagementType = (*CardServiceType2Code)(&value)
}

func (c *CardTransaction12) SetInitiatorDateTime(value string) {
	c.InitiatorDateTime = (*ISODateTime)(&value)
}

func (c *CardTransaction12) SetNumberOfMessages(value string) {
	c.NumberOfMessages = (*Number)(&value)
}

func (c *CardTransaction12) SetMaximumNumberOfMessages(value string) {
	c.MaximumNumberOfMessages = (*Number)(&value)
}

func (c *CardTransaction12) AddTransactionResponse() *ResponseType2 {
	c.TransactionResponse = new(ResponseType2)
	return c.TransactionResponse
}
