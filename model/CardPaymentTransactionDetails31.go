package model

// Details of the transaction in the authorisation request in a batch.
type CardPaymentTransactionDetails31 struct {

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Total amount of the transaction.
	TotalAmount *ImpliedCurrencyAndAmount `xml:"TtlAmt"`

	// Qualifies the amount associated with the transaction.
	AmountQualifier *TypeOfAmount1Code `xml:"AmtQlfr,omitempty"`

	// Detailed amounts associated with the total amount of transaction.
	DetailedAmount *DetailedAmount15 `xml:"DtldAmt,omitempty"`

	// Transaction authorisation deadline to complete the related payment.
	ValidityDate *ISODate `xml:"VldtyDt,omitempty"`

	// Reason to process an online authorisation.
	OnLineReason *OnLineReason1Code `xml:"OnLineRsn,omitempty"`

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

	// Codification used to identify the products.
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

func (c *CardPaymentTransactionDetails31) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransactionDetails31) SetTotalAmount(value, currency string) {
	c.TotalAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransactionDetails31) SetAmountQualifier(value string) {
	c.AmountQualifier = (*TypeOfAmount1Code)(&value)
}

func (c *CardPaymentTransactionDetails31) AddDetailedAmount() *DetailedAmount15 {
	c.DetailedAmount = new(DetailedAmount15)
	return c.DetailedAmount
}

func (c *CardPaymentTransactionDetails31) SetValidityDate(value string) {
	c.ValidityDate = (*ISODate)(&value)
}

func (c *CardPaymentTransactionDetails31) SetOnLineReason(value string) {
	c.OnLineReason = (*OnLineReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails31) SetUnattendedLevelCategory(value string) {
	c.UnattendedLevelCategory = (*Max35NumericText)(&value)
}

func (c *CardPaymentTransactionDetails31) SetAccountType(value string) {
	c.AccountType = (*CardAccountType3Code)(&value)
}

func (c *CardPaymentTransactionDetails31) AddCurrencyConversionResult() *CurrencyConversion8 {
	c.CurrencyConversionResult = new(CurrencyConversion8)
	return c.CurrencyConversionResult
}

func (c *CardPaymentTransactionDetails31) AddInstalment() *RecurringTransaction2 {
	c.Instalment = new(RecurringTransaction2)
	return c.Instalment
}

func (c *CardPaymentTransactionDetails31) AddAggregationTransaction() *AggregationTransaction2 {
	c.AggregationTransaction = new(AggregationTransaction2)
	return c.AggregationTransaction
}

func (c *CardPaymentTransactionDetails31) SetProductCodeSetIdentification(value string) {
	c.ProductCodeSetIdentification = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails31) AddSaleItem() *Product3 {
	newValue := new(Product3)
	c.SaleItem = append(c.SaleItem, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails31) SetDeliveryLocation(value string) {
	c.DeliveryLocation = (*Max10Text)(&value)
}

func (c *CardPaymentTransactionDetails31) AddCardPaymentInvoice() *CardPaymentInvoice2 {
	c.CardPaymentInvoice = new(CardPaymentInvoice2)
	return c.CardPaymentInvoice
}

func (c *CardPaymentTransactionDetails31) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max10000Binary)(&value)
}
