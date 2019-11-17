package model

// Details of the transaction in the authorisation response contained in a batch.
type CardPaymentTransactionDetails32 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *ImpliedCurrencyAndAmount `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails32) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails32) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails32) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails32) SetInvoiceAmount(value, currency string) {
	c.InvoiceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails32) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails32) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails32) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
