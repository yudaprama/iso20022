package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext17 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext17 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext17) AddPaymentContext() *PaymentContext17 {
	c.PaymentContext = new(PaymentContext17)
	return c.PaymentContext
}

func (c *CardPaymentContext17) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
