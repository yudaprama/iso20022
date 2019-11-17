package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext18 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext18 `xml:"PmtCntxt,omitempty"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext18) AddPaymentContext() *PaymentContext18 {
	c.PaymentContext = new(PaymentContext18)
	return c.PaymentContext
}

func (c *CardPaymentContext18) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
