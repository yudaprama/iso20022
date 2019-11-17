package model

// Details of the transaction to capture.
type CardPaymentTransactionDetails30 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Amount requested to be authorised.
	RequestedAmount *ImpliedCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Amount authorised for the transaction.
	AuthorisedAmount *ImpliedCurrencyAndAmount `xml:"AuthrsdAmt,omitempty"`

	// Amount of the transaction that will be invoiced to the cardholder.
	InvoiceAmount *ImpliedCurrencyAndAmount `xml:"InvcAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Transaction category level on an unattended POI (Point Of Interaction).
	UnattendedLevelCategory *Max35NumericText `xml:"UattnddLvlCtgy,omitempty"`

	// Type of cardholder account used for the transaction.
	AccountType *CardAccountType3Code `xml:"AcctTp,omitempty"`

	// Result of the currency conversion proposed to the cardholder.
	CurrencyConversionResult *CurrencyConversion8 `xml:"CcyConvsRslt,omitempty"`

	// Data related to a financial loan (instalment) or to a recurring transaction.
	Instalment *RecurringTransaction2 `xml:"Instlmt,omitempty"`

	// Payment transaction with an aggregated amount.
	AggregationTransaction *AggregationTransaction2 `xml:"AggtnTx,omitempty"`

	// Identification of the product codes that are purchased.
	ProductCodeSetIdentification *Max10Text `xml:"PdctCdSetId,omitempty"`

	// Item purchased with the transaction.
	SaleItem []*Product3 `xml:"SaleItm,omitempty"`

	// Location of the delivery, for instance pump number or parking bay.
	DeliveryLocation *Max10Text `xml:"DlvryLctn,omitempty"`

	// Detailed invoice data.
	CardPaymentInvoice *CardPaymentInvoice2 `xml:"CardPmtInvc,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max10000Binary `xml:"ICCRltdData,omitempty"`
}

func (c *CardPaymentTransactionDetails30) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails30) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails30) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails30) SetRequestedAmount(value, currency string) {
	c.RequestedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetAuthorisedAmount(value, currency string) {
	c.AuthorisedAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetInvoiceAmount(value, currency string) {
	c.InvoiceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails30) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails30) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails30) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails30) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails30) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails30) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails30) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails30) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails30) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails30) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails30) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
