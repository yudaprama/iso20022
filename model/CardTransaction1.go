package model

// Card transaction details.
type CardTransaction1 struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	Card *PaymentCard4 `xml:"Card,omitempty"`

	// Physical or logical card payment terminal containing software and hardware components.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`

	// Card transaction details, which can be either globalised by the acquirer or individual transaction.
	Transaction *CardTransaction1Choice `xml:"Tx,omitempty"`
}

func (c *CardTransaction1) AddCard() *PaymentCard4 {
	c.Card = new(PaymentCard4)
	return c.Card
}

func (c *CardTransaction1) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardTransaction1) AddTransaction() *CardTransaction1Choice {
	c.Transaction = new(CardTransaction1Choice)
	return c.Transaction
}
