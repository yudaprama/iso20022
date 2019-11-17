package model

// Details of the card transaction.
type CardTransactionDetail6 struct {

	// Amounts of the transaction expressed within the terminal currency.
	TransactionAmounts *CardTransactionAmount5 `xml:"TxAmts"`

	// Fees between acquirer and issuer exclusive of the transaction amount, and expressed in the currency of the reconciliation.
	TransactionFees []*DetailedAmount11 `xml:"TxFees,omitempty"`

	// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
	AdditionalAmounts []*DetailedAmount10 `xml:"AddtlAmts,omitempty"`

	// Data related to an integrated circuit card application.
	// It corresponds to ISO 8583, field number 55 for the versions 93 and 2003.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardTransactionDetail6) AddTransactionAmounts() *CardTransactionAmount5 {
	c.TransactionAmounts = new(CardTransactionAmount5)
	return c.TransactionAmounts
}

func (c *CardTransactionDetail6) AddTransactionFees() *DetailedAmount11 {
	newValue := new(DetailedAmount11)
	c.TransactionFees = append(c.TransactionFees, newValue)
	return newValue
}

func (c *CardTransactionDetail6) AddAdditionalAmounts() *DetailedAmount10 {
	newValue := new(DetailedAmount10)
	c.AdditionalAmounts = append(c.AdditionalAmounts, newValue)
	return newValue
}

func (c *CardTransactionDetail6) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
