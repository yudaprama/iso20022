package model

// Key exchange transaction.
type CardTransaction13 struct {

	// Type of key exchange.
	KeyExchangeType *CardServiceType3Code `xml:"KeyXchgTp"`

	// Date and time of the transaction.
	InitiatorDateTime *ISODateTime `xml:"InitrDtTm"`

	// Key that must be created and sent in the response, or that must be verified..
	RequestedKey []*KEKIdentifier3 `xml:"ReqdKey,omitempty"`

	// Created key to be stored.
	Key []*CryptographicKey6 `xml:"Key,omitempty"`

	// Response to the key exchange request.
	TransactionResponse *ResponseType2 `xml:"TxRspn,omitempty"`
}

func (c *CardTransaction13) SetKeyExchangeType(value string) {
	c.KeyExchangeType = (*CardServiceType3Code)(&value)
}

func (c *CardTransaction13) SetInitiatorDateTime(value string) {
	c.InitiatorDateTime = (*ISODateTime)(&value)
}

func (c *CardTransaction13) AddRequestedKey() *KEKIdentifier3 {
	newValue := new(KEKIdentifier3)
	c.RequestedKey = append(c.RequestedKey, newValue)
	return newValue
}

func (c *CardTransaction13) AddKey() *CryptographicKey6 {
	newValue := new(CryptographicKey6)
	c.Key = append(c.Key, newValue)
	return newValue
}

func (c *CardTransaction13) AddTransactionResponse() *ResponseType2 {
	c.TransactionResponse = new(ResponseType2)
	return c.TransactionResponse
}
