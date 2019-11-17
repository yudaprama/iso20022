package model

// Details of the transaction in the completion advice.
type CardPaymentTransactionDetails14 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount5 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the payment transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType1Code `xml:"AcctTp,omitempty"`

	// Indicates if a currency conversion has been accepted by the cardholder.
	ConversionAccepted *TrueFalseIndicator `xml:"ConvsAccptd,omitempty"`

	// Currency conversion proposed to the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Product purchased with the transaction.
	Product []*Product1 `xml:"Pdct,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice1 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails14) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails14) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails14) AddDetailedAmount() *DetailedAmount5 {
	c.DetailedAmount = new(DetailedAmount5)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails14) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails14) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails14) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails14) SetAccountType(value string) {
	c.AccountType = (*CardAccountType1Code)(&value)
}

func (c *CardPaymentTransactionDetails14) SetConversionAccepted(value string) {
	c.ConversionAccepted = (*TrueFalseIndicator)(&value)
}

func (c *CardPaymentTransactionDetails14) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}

func (c *CardPaymentTransactionDetails14) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails14) AddProduct() *Product1 {
	newValue := new(Product1)
	c.Product = append(c.Product, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails14) AddCardPaymentInvoice() *CardPaymentInvoice1 {
	c.CardPaymentInvoice = new(CardPaymentInvoice1)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails14) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
