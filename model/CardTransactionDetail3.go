package model

// Details of the card transaction.
type CardTransactionDetail3 struct {

	// Amounts of the transaction expressed within the terminal currency.
	TransactionAmounts *CardTransactionAmount3 `xml:"TxAmts"`

	// Fees between acquirer and issuer exclusive of the transaction amount, and expressed in the currency of the reconciliation.
	TransactionFees []*DetailedAmount11 `xml:"TxFees,omitempty"`

	// Additional amounts from the processor or the issuer without financial impacts on the transaction amount.
	AdditionalAmounts []*DetailedAmount10 `xml:"AddtlAmts,omitempty"`

	// Reason to send a card acquirer to issuer message.
	// It corresponds to ISO 8583 field number 25 for the version 93, and field number 9 for the version 2003.
	MessageReason *MessageReason1Code `xml:"MsgRsn,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	// It corresponds to ISO 8583 field number 57 for the version 93, and field number 3 for the version 2003.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended terminal.
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Way to identify a customer account or a relationship to its account affected for debits, inquiries and the “from” account for transfers.
	AccountFrom *CardAccount1 `xml:"AcctFr,omitempty"`

	// Way to identify a customer account or a relationship to its account affected for credits and the “to” account for transfers.
	AccountTo *CardAccount1 `xml:"AcctTo,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Information requested against money laundering for a transfer transaction.
	AntiMoneyLaundering *AntiMoneyLaundering1 `xml:"AML,omitempty"`

	// Data related to an integrated circuit card application.
	// It corresponds to ISO 8583 field number 55 for the versions 93 and 2003.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardTransactionDetail3) AddTransactionAmounts() *CardTransactionAmount3 {
	c.TransactionAmounts = new(CardTransactionAmount3)
	return c.TransactionAmounts
}

func (c *CardTransactionDetail3) AddTransactionFees() *DetailedAmount11 {
	newValue := new(DetailedAmount11)
	c.TransactionFees = append(c.TransactionFees, newValue)
	return newValue
}

func (c *CardTransactionDetail3) AddAdditionalAmounts() *DetailedAmount10 {
	newValue := new(DetailedAmount10)
	c.AdditionalAmounts = append(c.AdditionalAmounts, newValue)
	return newValue
}

func (c *CardTransactionDetail3) SetMessageReason(value string) {
	c.MessageReason = (*MessageReason1Code)(&value)
}

func (c *CardTransactionDetail3) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardTransactionDetail3) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardTransactionDetail3) AddAccountFrom() *CardAccount1 {
	c.AccountFrom = new(CardAccount1)
	return c.AccountFrom
}

func (c *CardTransactionDetail3) AddAccountTo() *CardAccount1 {
	c.AccountTo = new(CardAccount1)
	return c.AccountTo
}

func (c *CardTransactionDetail3) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardTransactionDetail3) AddAntiMoneyLaundering() *AntiMoneyLaundering1 {
	c.AntiMoneyLaundering = new(AntiMoneyLaundering1)
	return c.AntiMoneyLaundering
}

func (c *CardTransactionDetail3) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
