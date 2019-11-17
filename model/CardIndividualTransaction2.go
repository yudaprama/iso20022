package model

// Individual card transaction entry details.
type CardIndividualTransaction2 struct {

	// Data related to an integrated circuit card application.
	ICCRelatedData *Max1025Text `xml:"ICCRltdData,omitempty"`

	// Context of the card payment transaction.
	PaymentContext *PaymentContext3 `xml:"PmtCntxt,omitempty"`

	// Service in addition to the main service.
	AdditionalService *CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	// This element is also known as the MerchantCategoryCode.
	TransactionCategory *ExternalCardTransactionCategory1Code `xml:"TxCtgy,omitempty"`

	// Unique identification of the sales reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	SaleReconciliationIdentification *Max35Text `xml:"SaleRcncltnId,omitempty"`

	// Unique reference of the sales as provided by the merchant.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

	// Reason for representment of a card transaction.
	RePresentmentReason *ExternalRePresentmentReason1Code `xml:"RePresntmntRsn,omitempty"`

	// Sequential number of the card transaction, as assigned by the POI (Point of Interaction).
	// Usage: The sequential number is increased incrementally for each transaction.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId,omitempty"`

	// Product purchased with the transaction.
	Product *Product2 `xml:"Pdct,omitempty"`

	// Date when the deposit was validated by the financial institution that collected the cash.
	ValidationDate *ISODate `xml:"VldtnDt,omitempty"`

	// Sequential number of the validation of the cash deposit.
	// Usage: The sequential number is increased incrementally for each transaction.
	ValidationSequenceNumber *Max35Text `xml:"VldtnSeqNb,omitempty"`
}

func (c *CardIndividualTransaction2) SetICCRelatedData(value string) {
	c.ICCRelatedData = (*Max1025Text)(&value)
}

func (c *CardIndividualTransaction2) AddPaymentContext() *PaymentContext3 {
	c.PaymentContext = new(PaymentContext3)
	return c.PaymentContext
}

func (c *CardIndividualTransaction2) SetAdditionalService(value string) {
	c.AdditionalService = (*CardPaymentServiceType2Code)(&value)
}

func (c *CardIndividualTransaction2) SetTransactionCategory(value string) {
	c.TransactionCategory = (*ExternalCardTransactionCategory1Code)(&value)
}

func (c *CardIndividualTransaction2) SetSaleReconciliationIdentification(value string) {
	c.SaleReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction2) SetSaleReferenceNumber(value string) {
	c.SaleReferenceNumber = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction2) SetRePresentmentReason(value string) {
	c.RePresentmentReason = (*ExternalRePresentmentReason1Code)(&value)
}

func (c *CardIndividualTransaction2) SetSequenceNumber(value string) {
	c.SequenceNumber = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction2) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardIndividualTransaction2) AddProduct() *Product2 {
	c.Product = new(Product2)
	return c.Product
}

func (c *CardIndividualTransaction2) SetValidationDate(value string) {
	c.ValidationDate = (*ISODate)(&value)
}

func (c *CardIndividualTransaction2) SetValidationSequenceNumber(value string) {
	c.ValidationSequenceNumber = (*Max35Text)(&value)
}
