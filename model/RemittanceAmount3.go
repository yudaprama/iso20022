package model

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type RemittanceAmount3 struct {

	// Amount specified is the exact amount due and payable to the creditor.
	DuePayableAmount *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`

	// Amount of discount to be applied to the amount due and payable to the creditor.
	DiscountAppliedAmount []*DiscountAmountAndType1 `xml:"DscntApldAmt,omitempty"`

	// Amount of a credit note.
	CreditNoteAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`

	// Amount of the tax.
	TaxAmount []*TaxAmountAndType1 `xml:"TaxAmt,omitempty"`

	// Specifies detailed information on the amount and reason of the adjustment.
	AdjustmentAmountAndReason []*DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn,omitempty"`

	// Amount of money remitted.
	RemittedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

func (r *RemittanceAmount3) SetDuePayableAmount(value, currency string) {
	r.DuePayableAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount3) AddDiscountAppliedAmount() *DiscountAmountAndType1 {
	newValue := new(DiscountAmountAndType1)
	r.DiscountAppliedAmount = append(r.DiscountAppliedAmount, newValue)
	return newValue
}

func (r *RemittanceAmount3) SetCreditNoteAmount(value, currency string) {
	r.CreditNoteAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount3) AddTaxAmount() *TaxAmountAndType1 {
	newValue := new(TaxAmountAndType1)
	r.TaxAmount = append(r.TaxAmount, newValue)
	return newValue
}

func (r *RemittanceAmount3) AddAdjustmentAmountAndReason() *DocumentAdjustment1 {
	newValue := new(DocumentAdjustment1)
	r.AdjustmentAmountAndReason = append(r.AdjustmentAmountAndReason, newValue)
	return newValue
}

func (r *RemittanceAmount3) SetRemittedAmount(value, currency string) {
	r.RemittedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
