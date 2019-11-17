package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext2 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext2 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext2) AddPaymentContext() *PaymentContext2 {
	c.PaymentContext = new(PaymentContext2)
	return c.PaymentContext
}

func (c *CardPaymentContext2) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
