package model

// Environment of the transaction.
type CardPaymentEnvironment9 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard5 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder3 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment9) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment9) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment9) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment9) AddCard() *PaymentCard5 {
	c.Card = new(PaymentCard5)
	return c.Card
}

func (c *CardPaymentEnvironment9) AddCardholder() *Cardholder3 {
	c.Cardholder = new(Cardholder3)
	return c.Cardholder
}
