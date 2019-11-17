package model

// Specifies totals related to the invoice.
type InvoiceTotals1 struct {

	// Total amount subject to tax.
	TotalTaxableAmount *ActiveCurrencyAndAmount `xml:"TtlTaxblAmt"`

	// Sum of all tax amounts related to the invoice.
	TotalTaxAmount *ActiveCurrencyAndAmount `xml:"TtlTaxAmt"`

	// Variance on invoice amount taking into account discounts, allowances and charges.
	Adjustment *Adjustment5 `xml:"Adjstmnt,omitempty"`

	// Total amount of the invoice, being the sum of total invoice lines amounts, total invoice adjustment amount (discounts, allowances and charges) and total tax amounts.
	TotalInvoiceAmount *ActiveCurrencyAndAmount `xml:"TtlInvcAmt"`

	// Due date for the payment of the invoice.
	PaymentDueDate *ISODate `xml:"PmtDueDt"`
}

func (i *InvoiceTotals1) SetTotalTaxableAmount(value, currency string) {
	i.TotalTaxableAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvoiceTotals1) SetTotalTaxAmount(value, currency string) {
	i.TotalTaxAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvoiceTotals1) AddAdjustment() *Adjustment5 {
	i.Adjustment = new(Adjustment5)
	return i.Adjustment
}

func (i *InvoiceTotals1) SetTotalInvoiceAmount(value, currency string) {
	i.TotalInvoiceAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvoiceTotals1) SetPaymentDueDate(value string) {
	i.PaymentDueDate = (*ISODate)(&value)
}
