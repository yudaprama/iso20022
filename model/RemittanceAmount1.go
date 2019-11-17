package model

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type RemittanceAmount1 struct {

	// Amount specified is the exact amount due and payable to the creditor.
	DuePayableAmount *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`

	// Amount of money that results from the application of an agreed discount to the amount due and payable to the creditor.
	DiscountAppliedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"DscntApldAmt,omitempty"`

	// Amount specified for the referred document is the amount of a credit note.
	CreditNoteAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`

	// Quantity of cash resulting from the calculation of the tax.
	TaxAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TaxAmt,omitempty"`

	// Set of elements used to provide information on the amount and reason of the document adjustment.
	AdjustmentAmountAndReason []*DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn,omitempty"`

	// Amount of money remitted for the referred document.
	RemittedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

func (r *RemittanceAmount1) SetDuePayableAmount(value, currency string) {
	r.DuePayableAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount1) SetDiscountAppliedAmount(value, currency string) {
	r.DiscountAppliedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount1) SetCreditNoteAmount(value, currency string) {
	r.CreditNoteAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount1) SetTaxAmount(value, currency string) {
	r.TaxAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount1) AddAdjustmentAmountAndReason() *DocumentAdjustment1 {
	newValue := new(DocumentAdjustment1)
	r.AdjustmentAmountAndReason = append(r.AdjustmentAmountAndReason, newValue)
	return newValue
}

func (r *RemittanceAmount1) SetRemittedAmount(value, currency string) {
	r.RemittedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
