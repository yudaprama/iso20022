package model

// Details of the card transaction.
type CardTransactionDetail4 struct {

	// Amounts of the transaction expressed within the terminal currency.
	TransactionAmounts *CardTransactionAmount4 `xml:"TxAmts"`

	// Fees between acquirer and issuer exclusive of the transaction amount, and expressed in the currency of the reconciliation.
	TransactionFees []*DetailedAmount11 `xml:"TxFees,omitempty"`

	// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
	AdditionalAmounts []*DetailedAmount10 `xml:"AddtlAmts,omitempty"`

	// Account involved in the card transaction.
	AccountAndBalance []*CardAccount2 `xml:"AcctAndBal,omitempty"`

	// Results of the verifications performed by the various agents during the processing of the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	// It corresponds to ISO 8583, field number 57 for the version 93, and 3 for the version 2003.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Data related to an integrated circuit card application.
	// It corresponds to ISO 8583, field number 55 for the versions 93 and 2003.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardTransactionDetail4) AddTransactionAmounts() *CardTransactionAmount4 {
	c.TransactionAmounts = new(CardTransactionAmount4)
	return c.TransactionAmounts
}

func (c *CardTransactionDetail4) AddTransactionFees() *DetailedAmount11 {
	newValue := new(DetailedAmount11)
	c.TransactionFees = append(c.TransactionFees, newValue)
	return newValue
}

func (c *CardTransactionDetail4) AddAdditionalAmounts() *DetailedAmount10 {
	newValue := new(DetailedAmount10)
	c.AdditionalAmounts = append(c.AdditionalAmounts, newValue)
	return newValue
}

func (c *CardTransactionDetail4) AddAccountAndBalance() *CardAccount2 {
	newValue := new(CardAccount2)
	c.AccountAndBalance = append(c.AccountAndBalance, newValue)
	return newValue
}

func (c *CardTransactionDetail4) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardTransactionDetail4) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardTransactionDetail4) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
