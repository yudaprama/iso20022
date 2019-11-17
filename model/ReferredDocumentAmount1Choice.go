package model

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type ReferredDocumentAmount1Choice struct {

	// Amount specified is the exact amount due and payable to the creditor.
	DuePayableAmount *CurrencyAndAmount `xml:"DuePyblAmt"`

	// Amount of money that results from the application of an agreed discount to the amount due and payable to the creditor.
	DiscountAppliedAmount *CurrencyAndAmount `xml:"DscntApldAmt"`

	// Amount of money remitted for the referred document.
	RemittedAmount *CurrencyAndAmount `xml:"RmtdAmt"`

	// Amount specified for the referred document is the amount of a credit note.
	CreditNoteAmount *CurrencyAndAmount `xml:"CdtNoteAmt"`

	// Quantity of cash resulting from the calculation of the tax.
	TaxAmount *CurrencyAndAmount `xml:"TaxAmt"`
}

func (r *ReferredDocumentAmount1Choice) SetDuePayableAmount(value, currency string) {
	r.DuePayableAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReferredDocumentAmount1Choice) SetDiscountAppliedAmount(value, currency string) {
	r.DiscountAppliedAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReferredDocumentAmount1Choice) SetRemittedAmount(value, currency string) {
	r.RemittedAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReferredDocumentAmount1Choice) SetCreditNoteAmount(value, currency string) {
	r.CreditNoteAmount = NewCurrencyAndAmount(value, currency)
}

func (r *ReferredDocumentAmount1Choice) SetTaxAmount(value, currency string) {
	r.TaxAmount = NewCurrencyAndAmount(value, currency)
}
