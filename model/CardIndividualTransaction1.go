package model

// Individual card transaction entry details.
type CardIndividualTransaction1 struct {

	// Service in addition to the main service.
	AdditionalService *CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Category code conform to ISO 18245, related to the type of services or goods the merchant provides for the transaction.
	TransactionCategory *ExternalCardTransactionCategory1Code `xml:"TxCtgy,omitempty"`

	// Unique identification of the sales reconciliation period between the acceptor and the acquirer. This identification might be linked to the identification of the settlement for further verification by the merchant.
	SaleReconciliationIdentification *Max35Text `xml:"SaleRcncltnId,omitempty"`

	// Unique reference of the sales as provided by the merchant.
	SaleReferenceNumber *Max35Text `xml:"SaleRefNb,omitempty"`

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

func (c *CardIndividualTransaction1) SetAdditionalService(value string) {
	c.AdditionalService = (*CardPaymentServiceType2Code)(&value)
}

func (c *CardIndividualTransaction1) SetTransactionCategory(value string) {
	c.TransactionCategory = (*ExternalCardTransactionCategory1Code)(&value)
}

func (c *CardIndividualTransaction1) SetSaleReconciliationIdentification(value string) {
	c.SaleReconciliationIdentification = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction1) SetSaleReferenceNumber(value string) {
	c.SaleReferenceNumber = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction1) SetSequenceNumber(value string) {
	c.SequenceNumber = (*Max35Text)(&value)
}

func (c *CardIndividualTransaction1) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardIndividualTransaction1) AddProduct() *Product2 {
	c.Product = new(Product2)
	return c.Product
}

func (c *CardIndividualTransaction1) SetValidationDate(value string) {
	c.ValidationDate = (*ISODate)(&value)
}

func (c *CardIndividualTransaction1) SetValidationSequenceNumber(value string) {
	c.ValidationSequenceNumber = (*Max35Text)(&value)
}
