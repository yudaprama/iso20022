package model

// Details of the card transaction.
type CardTransactionDetail5 struct {

	// Amounts of the transaction expressed within the terminal currency.
	TransactionAmounts *CardTransactionAmount5 `xml:"TxAmts"`

	// Fees between acquirer and issuer exclusive of the transaction amount, and expressed in the currency of the reconciliation.
	TransactionFees []*DetailedAmount11 `xml:"TxFees,omitempty"`

	// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
	AdditionalAmounts []*DetailedAmount10 `xml:"AddtlAmts,omitempty"`

	// Reason to send a card acquirer to issuer message.
	// It corresponds to ISO 8583, field number 25 for the version 93, and 9 for the version 2003.
	MessageReason []*MessageReason1Code `xml:"MsgRsn"`

	// Data related to an integrated circuit card application.
	// It corresponds to ISO 8583, field number 55 for the versions 93 and 2003.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardTransactionDetail5) AddTransactionAmounts() *CardTransactionAmount5 {
	c.TransactionAmounts = new(CardTransactionAmount5)
	return c.TransactionAmounts
}

func (c *CardTransactionDetail5) AddTransactionFees() *DetailedAmount11 {
	newValue := new(DetailedAmount11)
	c.TransactionFees = append(c.TransactionFees, newValue)
	return newValue
}

func (c *CardTransactionDetail5) AddAdditionalAmounts() *DetailedAmount10 {
	newValue := new(DetailedAmount10)
	c.AdditionalAmounts = append(c.AdditionalAmounts, newValue)
	return newValue
}

func (c *CardTransactionDetail5) AddMessageReason(value string) {
	c.MessageReason = append(c.MessageReason, (*MessageReason1Code)(&value))
}

func (c *CardTransactionDetail5) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
