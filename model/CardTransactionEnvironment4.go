package model

// Environment of the transaction.
type CardTransactionEnvironment4 struct {

	// Acquirer identification of the card transaction.
	AcquirerIdentification *Max35Text `xml:"AcqrrId"`

	// Identification of the interconnected card scheme from which the request is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Identification of the card acceptor performing the transaction.
	AcceptorIdentification *Max35Text `xml:"AccptrId,omitempty"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *Max35Text `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard15 `xml:"Card"`

	// Payment token information.
	PaymentToken *CardPaymentToken2 `xml:"PmtTkn,omitempty"`
}

func (c *CardTransactionEnvironment4) SetAcquirerIdentification(value string) {
	c.AcquirerIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetAcceptorIdentification(value string) {
	c.AcceptorIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) SetTerminalIdentification(value string) {
	c.TerminalIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment4) AddCard() *PaymentCard15 {
	c.Card = new(PaymentCard15)
	return c.Card
}

func (c *CardTransactionEnvironment4) AddPaymentToken() *CardPaymentToken2 {
	c.PaymentToken = new(CardPaymentToken2)
	return c.PaymentToken
}
