package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext6 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext6 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext6) AddPaymentContext() *PaymentContext6 {
	c.PaymentContext = new(PaymentContext6)
	return c.PaymentContext
}

func (c *CardPaymentContext6) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
