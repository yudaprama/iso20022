package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext16 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext16 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext2 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext16) AddPaymentContext() *PaymentContext16 {
	c.PaymentContext = new(PaymentContext16)
	return c.PaymentContext
}

func (c *CardPaymentContext16) AddSaleContext() *SaleContext2 {
	c.SaleContext = new(SaleContext2)
	return c.SaleContext
}
