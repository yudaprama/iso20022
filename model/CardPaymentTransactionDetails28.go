package model

// Details of the transaction in the authorisation response.
type CardPaymentTransactionDetails28 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *DetailedAmount4 `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails28) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails28) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails28) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails28) AddInvoiceAmount() *DetailedAmount4 {
	c.InvoiceAmount = new(DetailedAmount4)
	return c.InvoiceAmount
}

func (c *CardPaymentTransactionDetails28) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails28) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails28) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
