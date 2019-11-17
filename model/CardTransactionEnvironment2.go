package model

// Environment of the transaction.
type CardTransactionEnvironment2 struct {

	// Acquirer identification of the transaction.
	AcquirerIdentification *Max35Text `xml:"AcqrrId"`

	// Identification of the interconnected card scheme from which the response is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Identification of the card acceptor performing the transaction.
	AcceptorIdentification *Max35Text `xml:"AccptrId,omitempty"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *Max35Text `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard13 `xml:"Card"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`
}

func (c *CardTransactionEnvironment2) SetAcquirerIdentification(value string) {
	c.AcquirerIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetAcceptorIdentification(value string) {
	c.AcceptorIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) SetTerminalIdentification(value string) {
	c.TerminalIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment2) AddCard() *PaymentCard13 {
	c.Card = new(PaymentCard13)
	return c.Card
}

func (c *CardTransactionEnvironment2) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}

func (c *CardTransactionEnvironment2) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}
