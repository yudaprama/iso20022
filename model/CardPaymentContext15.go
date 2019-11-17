package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext15 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext15 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext15) AddPaymentContext() *PaymentContext15 {
	c.PaymentContext = new(PaymentContext15)
	return c.PaymentContext
}

func (c *CardPaymentContext15) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
