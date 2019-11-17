package model

// Network management transaction.
type CardTransaction11 struct {

	// Type of network management service (correspond to the ISO 8583 field 24)
	NetworkManagementType *CardServiceType2Code `xml:"NtwkMgmtTp"`

	// Number of messages in the store and forward queue.
	NumberOfMessages *Number `xml:"NbOfMsgs,omitempty"`

	// Maximum number of messages in the store and forward queue.
	MaximumNumberOfMessages *Number `xml:"MaxNbOfMsgs,omitempty"`

	// Date and time of the transaction.
	InitiatorDateTime *ISODateTime `xml:"InitrDtTm,omitempty"`

	// Response to the network management request.
	TransactionResponse *ResponseType2 `xml:"TxRspn,omitempty"`
}

func (c *CardTransaction11) SetNetworkManagementType(value string) {
	c.NetworkManagementType = (*CardServiceType2Code)(&value)
}

func (c *CardTransaction11) SetNumberOfMessages(value string) {
	c.NumberOfMessages = (*Number)(&value)
}

func (c *CardTransaction11) SetMaximumNumberOfMessages(value string) {
	c.MaximumNumberOfMessages = (*Number)(&value)
}

func (c *CardTransaction11) SetInitiatorDateTime(value string) {
	c.InitiatorDateTime = (*ISODateTime)(&value)
}

func (c *CardTransaction11) AddTransactionResponse() *ResponseType2 {
	c.TransactionResponse = new(ResponseType2)
	return c.TransactionResponse
}
