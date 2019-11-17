package model

// Environment of the transaction.
type CardPaymentEnvironment54 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard19 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment54) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment54) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment54) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment54) AddCard() *PaymentCard19 {
	c.Card = new(PaymentCard19)
	return c.Card
}

func (c *CardPaymentEnvironment54) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}
