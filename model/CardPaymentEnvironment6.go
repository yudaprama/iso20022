package model

// Environment of the card payment transaction.
type CardPaymentEnvironment6 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard3 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder2 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment6) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment6) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment6) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment6) AddCard() *PaymentCard3 {
	c.Card = new(PaymentCard3)
	return c.Card
}

func (c *CardPaymentEnvironment6) AddCardholder() *Cardholder2 {
	c.Cardholder = new(Cardholder2)
	return c.Cardholder
}
