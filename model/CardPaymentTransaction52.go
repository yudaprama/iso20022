package model

// Identification of the original transaction.
type CardPaymentTransaction52 struct {

	// Global reference of the sale transaction for the sale system.
	SaleReferenceIdentification *Max35Text `xml:"SaleRefId,omitempty"`

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType5Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType9Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Entry mode of the card information.
	CardDataEntryMode *CardDataReading5Code `xml:"CardDataNtryMd,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult3 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction52) SetSaleReferenceIdentification(value string) {
	c.SaleReferenceIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction52) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction52) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction52) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType5Code)(&value)
}

func (c *CardPaymentTransaction52) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType9Code)(&value))
}

func (c *CardPaymentTransaction52) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction52) SetCardDataEntryMode(value string) {
	c.CardDataEntryMode = (*CardDataReading5Code)(&value)
}

func (c *CardPaymentTransaction52) AddTransactionResult() *CardPaymentTransactionResult3 {
	c.TransactionResult = new(CardPaymentTransactionResult3)
	return c.TransactionResult
}
