package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext14 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext14 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext14) AddPaymentContext() *PaymentContext14 {
	c.PaymentContext = new(PaymentContext14)
	return c.PaymentContext
}

func (c *CardPaymentContext14) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
