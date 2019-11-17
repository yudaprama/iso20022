package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext7 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext7 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext7) AddPaymentContext() *PaymentContext7 {
	c.PaymentContext = new(PaymentContext7)
	return c.PaymentContext
}

func (c *CardPaymentContext7) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
