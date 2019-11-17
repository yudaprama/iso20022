package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext12 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext12 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext12) AddPaymentContext() *PaymentContext12 {
	c.PaymentContext = new(PaymentContext12)
	return c.PaymentContext
}

func (c *CardPaymentContext12) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
