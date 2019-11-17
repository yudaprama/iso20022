package model

// Environment of the transaction given in a response to a request.
type CardPaymentEnvironment33 struct {

	// Acquirer involved in the card payment.
	AcquirerIdentification *GenericIdentification53 `xml:"AcqrrId,omitempty"`

	// Identification of the merchant.
	MerchantIdentification *GenericIdentification32 `xml:"MrchntId,omitempty"`

	// Identification of the POI (Point Of Interaction) performing the transaction.
	POIIdentification *GenericIdentification32 `xml:"POIId"`

	// Card performing the transaction.
	Card *PaymentCard10 `xml:"Card,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardPaymentEnvironment33) AddAcquirerIdentification() *GenericIdentification53 {
	c.AcquirerIdentification = new(GenericIdentification53)
	return c.AcquirerIdentification
}

func (c *CardPaymentEnvironment33) AddMerchantIdentification() *GenericIdentification32 {
	c.MerchantIdentification = new(GenericIdentification32)
	return c.MerchantIdentification
}

func (c *CardPaymentEnvironment33) AddPOIIdentification() *GenericIdentification32 {
	c.POIIdentification = new(GenericIdentification32)
	return c.POIIdentification
}

func (c *CardPaymentEnvironment33) AddCard() *PaymentCard10 {
	c.Card = new(PaymentCard10)
	return c.Card
}

func (c *CardPaymentEnvironment33) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}
