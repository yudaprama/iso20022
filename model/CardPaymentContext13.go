package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext13 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext13 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext13) AddPaymentContext() *PaymentContext13 {
	c.PaymentContext = new(PaymentContext13)
	return c.PaymentContext
}

func (c *CardPaymentContext13) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
