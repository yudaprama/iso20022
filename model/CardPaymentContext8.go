package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext8 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext8 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext8) AddPaymentContext() *PaymentContext8 {
	c.PaymentContext = new(PaymentContext8)
	return c.PaymentContext
}

func (c *CardPaymentContext8) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
