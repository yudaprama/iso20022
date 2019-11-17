package model

// Identification of the original transaction.
type CardPaymentTransaction8 struct {

	// Unique identification of the transaction assigned by the POI (Point Of Interaction).
	TransactionIdentification *TransactionIdentifier1 `xml:"TxId"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Identification of the transaction assigned by the initiating party for the recipient party.
	InitiatorTransactionIdentification *Max35Text `xml:"InitrTxId,omitempty"`

	// Identification of the transaction assigned by the recipient party for the initiating party.
	RecipientTransactionIdentification *Max35Text `xml:"RcptTxId,omitempty"`

	// Type of transaction being undertaken for the main service.
	TransactionType *CardPaymentServiceType1Code `xml:"TxTp"`

	// Service in addition to the main service.
	AdditionalService []*CardPaymentServiceType2Code `xml:"AddtlSvc,omitempty"`

	// Additional attribute of the service type.
	ServiceAttribute *CardPaymentServiceType3Code `xml:"SvcAttr,omitempty"`

	// Result of the original transaction.
	TransactionResult *CardPaymentTransactionResult1 `xml:"TxRslt,omitempty"`
}

func (c *CardPaymentTransaction8) AddTransactionIdentification() *TransactionIdentifier1 {
	c.TransactionIdentification = new(TransactionIdentifier1)
	return c.TransactionIdentification
}

func (c *CardPaymentTransaction8) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentTransaction8) SetInitiatorTransactionIdentification(value string) {
	c.InitiatorTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction8) SetRecipientTransactionIdentification(value string) {
	c.RecipientTransactionIdentification = (*Max35Text)(&value)
}

func (c *CardPaymentTransaction8) SetTransactionType(value string) {
	c.TransactionType = (*CardPaymentServiceType1Code)(&value)
}

func (c *CardPaymentTransaction8) AddAdditionalService(value string) {
	c.AdditionalService = append(c.AdditionalService, (*CardPaymentServiceType2Code)(&value))
}

func (c *CardPaymentTransaction8) SetServiceAttribute(value string) {
	c.ServiceAttribute = (*CardPaymentServiceType3Code)(&value)
}

func (c *CardPaymentTransaction8) AddTransactionResult() *CardPaymentTransactionResult1 {
	c.TransactionResult = new(CardPaymentTransactionResult1)
	return c.TransactionResult
}
