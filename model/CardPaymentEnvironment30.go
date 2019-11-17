package model

// Environment of the transaction.
type CardPaymentEnvironment30 struct {

	// Acquirer involved in the card payment transaction.
	AcquirerIdentification *GenericIdentification32 `xml:"AcqrrId,omitempty"`

	// Merchant involved in the card payment transaction.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard7 `xml:"Card"`

	// Language selected for the cardholder interface during the transaction.
	CardholderLanguage *ISO2ALanguageCode `xml:"CrdhldrLang,omitempty"`
}

func (c *CardPaymentEnvironment30) AddAcquirerIdentification() *GenericIdentification32 {
	c.AcquirerIdentification = new(GenericIdentification32)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment30) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment30) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment30) AddCard() *PaymentCard7 {
	c.Card = new(PaymentCard7)
	return c.Card
}

func (c *CardPaymentEnvironment30) SetCardholderLanguage(value string) {
	c.CardholderLanguage = (*ISO2ALanguageCode)(&value)
}
