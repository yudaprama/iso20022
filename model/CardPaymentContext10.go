package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext10 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext10 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext10) AddPaymentContext() *PaymentContext10 {
	c.PaymentContext = new(PaymentContext10)
	return c.PaymentContext
}

func (c *CardPaymentContext10) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
