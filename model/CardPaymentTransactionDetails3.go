package model

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails3 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount []*DetailedAmount1 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	RecurringTransaction *RecurringTransaction1 `xml:"RcrngTx,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails3) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails3) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails3) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails3) AddDetailedAmount() *DetailedAmount1 {
	newValue := new(DetailedAmount1)
	c.DetailedAmount = append(c.DetailedAmount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails3) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails3) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails3) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails3) AddRecurringTransaction() *RecurringTransaction1 {
	c.RecurringTransaction = new(RecurringTransaction1)
	return c.RecurringTransaction
}

func (c *CardPaymentTransactionDetails3) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails3) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
