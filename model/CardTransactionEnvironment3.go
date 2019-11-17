package model

// Environment of the transaction.
type CardTransactionEnvironment3 struct {

	// Acquirer of the card transaction.
	Acquirer *Acquirer6 `xml:"Acqrr"`

	// Identification of the interconnected card scheme from which the response is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Acceptor performing the card transaction.
	Acceptor *Organisation19 `xml:"Accptr"`

	// Identification of the card terminal performing the transaction.
	TerminalIdentification *GenericIdentification32 `xml:"TermnlId,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard14 `xml:"Card"`

	// Container of tenders used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken4 `xml:"PmtTkn,omitempty"`
}

func (c *CardTransactionEnvironment3) AddAcquirer() *Acquirer6 {
	c.Acquirer = new(Acquirer6)
	return c.Acquirer
}

func (c *CardTransactionEnvironment3) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment3) AddAcceptor() *Organisation19 {
	c.Acceptor = new(Organisation19)
	return c.Acceptor
}

func (c *CardTransactionEnvironment3) AddTerminalIdentification() *GenericIdentification32 {
	c.TerminalIdentification = new(GenericIdentification32)
	return c.TerminalIdentification
}

func (c *CardTransactionEnvironment3) AddCard() *PaymentCard14 {
	c.Card = new(PaymentCard14)
	return c.Card
}

func (c *CardTransactionEnvironment3) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardTransactionEnvironment3) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardTransactionEnvironment3) AddPaymentToken() *CardPaymentToken4 {
	c.PaymentToken = new(CardPaymentToken4)
	return c.PaymentToken
}
