package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext1 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext1 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext1) AddPaymentContext() *PaymentContext1 {
	c.PaymentContext = new(PaymentContext1)
	return c.PaymentContext
}

func (c *CardPaymentContext1) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
