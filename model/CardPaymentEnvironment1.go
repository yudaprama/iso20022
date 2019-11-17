package model

// Environment of the transaction.
type CardPaymentEnvironment1 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard1 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder1 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment1) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment1) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment1) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardPaymentEnvironment1) AddCard() *PaymentCard1 {
	c.Card = new(PaymentCard1)
	return c.Card
}

func (c *CardPaymentEnvironment1) AddCardholder() *Cardholder1 {
	c.Cardholder = new(Cardholder1)
	return c.Cardholder
}
