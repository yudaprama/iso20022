package model

// Details of the transaction in the authorisation request.
type CardPaymentTransactionDetails8 struct {

	// Amounts associated with the total amount of transaction.
	Amount []*CardAmountAndCurrencyExchange1 `xml:"Amt,omitempty"`

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max1025Text `xml:"ICCRltdData,omitempty"`

	// Context of the card payment transaction.
	PaymentContext *PaymentContext3 `xml:"PmtCntxt,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	MerchantCategoryCode *Min3Max4Text `xml:"MrchntCtgyCd,omitempty"`

	// Local date and time of the transaction assigned by the POI (Point Of Interaction).
	TransactionDateTime *ISODateTime `xml:"TxDtTm,omitempty"`

	// Identification of a sale transaction assigned by the sale system.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

	// Reason for representment of a card transaction.
	RePresentmentReason *ExternalRePresentmentReason1Code `xml:"RePresntmntRsn,omitempty"`

	// Service in addition to the main service.
	AdditionalService *CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Identification of the transaction that has to be unique for a time period.
	TransactionReference *Max35Text `xml:"TxRef,omitempty"`
}

func (c *CardPaymentTransactionDetails8) AddAmount() *CardAmountAndCurrencyExchange1 {
	newValue := new(CardAmountAndCurrencyExchange1)
	c.Amount = append(c.Amount, newValue)
	return newValue
}

func (c *CardPaymentTransactionDetails8) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max1025Text)(&value)
}

func (c *CardPaymentTransactionDetails8) AddPaymentContext() *PaymentContext3 {
	c.PaymentContext = new(PaymentContext3)
	return c.PaymentContext
}

func (c *CardPaymentTransactionDetails8) SetMerchantCategoryCode(value string) {
	c.MerchantCategoryCode = (*Min3Max4Text)(&value)
}

func (c *CardPaymentTransactionDetails8) SetTransactionDateTime(value string) {
	c.TransactionDateTime = (*ISODateTime)(&value)
}

func (c *CardPaymentTransactionDetails8) SetSaleReferenceNumber(value string) {
	c.SaleReferenceNumber = (*Max35Text)(&value)
}

func (c *CardPaymentTransactionDetails8) SetRePresentmentReason(value string) {
	c.RePresentmentReason = (*ExternalRePresentmentReason1Code)(&value)
}

func (c *CardPaymentTransactionDetails8) SetAdditionalService(value string) {
	c.AdditionalService = (*CardPaymentServiceType2Code)(&value)
}

func (c *CardPaymentTransactionDetails8) SetTransactionReference(value string) {
	c.TransactionReference = (*Max35Text)(&value)
}
