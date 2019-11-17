package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext4 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext4 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext4) AddPaymentContext() *PaymentContext4 {
	c.PaymentContext = new(PaymentContext4)
	return c.PaymentContext
}

func (c *CardPaymentContext4) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
