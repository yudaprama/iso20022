package model

// Environment of the cancellation.
type CardPaymentEnvironment4 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment cancellation.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard3 `xml:"Card"`
}

func (c *CardPaymentEnvironment4) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment4) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment4) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment4) AddCard() *PaymentCard3 {
	c.Card = new(PaymentCard3)
	return c.Card
}
