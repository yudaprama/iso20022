package model

// Key exchange transaction.
type CardTransaction14 struct {

	// Type of key exchange.
	KeyExchangeType *CardServiceType3Code `xml:"KeyXchgTp"`

	// Date and time of the transaction.
	InitiatorDateTime *ISODateTime `xml:"InitrDtTm"`

	// Key that has been verified.
	KeyVerification []*KEKIdentifier3 `xml:"KeyVrfctn,omitempty"`

	// Created key to be stored.
	Key []*CryptographicKey6 `xml:"Key,omitempty"`

	// Response to the key exchange request.
	TransactionResponse *ResponseType2 `xml:"TxRspn"`
}

func (c *CardTransaction14) SetKeyExchangeType(value string) {
	c.KeyExchangeType = (*CardServiceType3Code)(&value)
}

func (c *CardTransaction14) SetInitiatorDateTime(value string) {
	c.InitiatorDateTime = (*ISODateTime)(&value)
}

func (c *CardTransaction14) AddKeyVerification() *KEKIdentifier3 {
	newValue := new(KEKIdentifier3)
	c.KeyVerification = append(c.KeyVerification, newValue)
	return newValue
}

func (c *CardTransaction14) AddKey() *CryptographicKey6 {
	newValue := new(CryptographicKey6)
	c.Key = append(c.Key, newValue)
	return newValue
}

func (c *CardTransaction14) AddTransactionResponse() *ResponseType2 {
	c.TransactionResponse = new(ResponseType2)
	return c.TransactionResponse
}
