package model

// Environment of the transaction.
type CardPaymentEnvironment44 struct {

	// Acquirer involved in the card payment transaction.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Merchant involved in the card payment transaction.
	MerchantIdentification *GenericIdentification53 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard9 `xml:"Card"`

	// Language selected for the cardholder interface during the transaction.
	// Reference: ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	CardholderLanguage *LanguageCode `xml:"CrdhldrLang,omitempty"`
}

func (c *CardPaymentEnvironment44) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment44) AddMerchantIdentification() *GenericIdentification53 {
	c.MerchantIdentification = new(GenericIdentification53)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment44) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}

func (c *CardPaymentEnvironment44) AddCard() *PaymentCard9 {
	c.Card = new(PaymentCard9)
	return c.Card
}

func (c *CardPaymentEnvironment44) SetCardholderLanguage(value string) {
	c.CardholderLanguage = (*LanguageCode)(&value)
}
