package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext19 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext19 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext19) AddPaymentContext() *PaymentContext19 {
	c.PaymentContext = new(PaymentContext19)
	return c.PaymentContext
}

func (c *CardPaymentContext19) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
