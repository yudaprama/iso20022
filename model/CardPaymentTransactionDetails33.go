package model

// Details of the payment transaction.
type CardPaymentTransactionDetails33 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails33) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails33) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails33) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails33) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails33) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
