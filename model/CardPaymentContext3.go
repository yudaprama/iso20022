package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext3 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext1 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext3) AddPaymentContext() *PaymentContext1 {
	c.PaymentContext = new(PaymentContext1)
	return c.PaymentContext
}

func (c *CardPaymentContext3) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
