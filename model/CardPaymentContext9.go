package model

// Context in which the transaction is performed (payment and sale).
type CardPaymentContext9 struct {

	// Context of the card payment transaction.
	PaymentContext *PaymentContext9 `xml:"PmtCntxt"`

	// Context of the sale involving the card payment transaction.
	SaleContext *SaleContext1 `xml:"SaleCntxt,omitempty"`
}

func (c *CardPaymentContext9) AddPaymentContext() *PaymentContext9 {
	c.PaymentContext = new(PaymentContext9)
	return c.PaymentContext
}

func (c *CardPaymentContext9) AddSaleContext() *SaleContext1 {
	c.SaleContext = new(SaleContext1)
	return c.SaleContext
}
