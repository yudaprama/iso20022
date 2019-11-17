package model

// Card transaction details.
type CardTransaction2 struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	Card *PaymentCard4 `xml:"Card,omitempty"`

	// Physical or logical card payment terminal containing software and hardware components.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`

	// Card transaction details, which can be either globalised by the acquirer or individual transaction.
	Transaction *CardTransaction2Choice `xml:"Tx,omitempty"`

	// Prepaid account for the transfer or loading of an amount of money.
	PrePaidAccount *CashAccount24 `xml:"PrePdAcct,omitempty"`
}

func (c *CardTransaction2) AddCard() *PaymentCard4 {
	c.Card = new(PaymentCard4)
	return c.Card
}

func (c *CardTransaction2) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardTransaction2) AddTransaction() *CardTransaction2Choice {
	c.Transaction = new(CardTransaction2Choice)
	return c.Transaction
}

func (c *CardTransaction2) AddPrePaidAccount() *CashAccount24 {
	c.PrePaidAccount = new(CashAccount24)
	return c.PrePaidAccount
}
