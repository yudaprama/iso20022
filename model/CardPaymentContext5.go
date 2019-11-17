package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext5 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext5 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext5) AddPaymentContext() *PaymentContext5 {
	c.PaymentContext = new(PaymentContext5)
	return c.PaymentContext
}

func (c *CardPaymentContext5) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
