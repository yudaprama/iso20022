package model

// Nature of the amount and currency on a document referred to in the remittance section, typically either the original amount due/payable or the amount actually remitted for the referenced document.
type RemittanceAmount2 struct {

	// Amount specified is the exact amount due and payable to the creditor.
	DuePayableAmount *ActiveOrHistoricCurrencyAndAmount `xml:"DuePyblAmt,omitempty"`

	// Amount specified for the referred document is the amount of discount to be applied to the amount due and payable to the creditor.
	DiscountAppliedAmount []*DiscountAmountAndType1 `xml:"DscntApldAmt,omitempty"`

	// Amount specified for the referred document is the amount of a credit note.
	CreditNoteAmount *ActiveOrHistoricCurrencyAndAmount `xml:"CdtNoteAmt,omitempty"`

	// Quantity of cash resulting from the calculation of the tax.
	TaxAmount []*TaxAmountAndType1 `xml:"TaxAmt,omitempty"`

	// Specifies detailed information on the amount and reason of the document adjustment.
	AdjustmentAmountAndReason []*DocumentAdjustment1 `xml:"AdjstmntAmtAndRsn,omitempty"`

	// Amount of money remitted for the referred document.
	RemittedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"RmtdAmt,omitempty"`
}

func (r *RemittanceAmount2) SetDuePayableAmount(value, currency string) {
	r.DuePayableAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount2) AddDiscountAppliedAmount() *DiscountAmountAndType1 {
	newValue := new(DiscountAmountAndType1)
	r.DiscountAppliedAmount = append(r.DiscountAppliedAmount, newValue)
	return newValue
}

func (r *RemittanceAmount2) SetCreditNoteAmount(value, currency string) {
	r.CreditNoteAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (r *RemittanceAmount2) AddTaxAmount() *TaxAmountAndType1 {
	newValue := new(TaxAmountAndType1)
	r.TaxAmount = append(r.TaxAmount, newValue)
	return newValue
}

func (r *RemittanceAmount2) AddAdjustmentAmountAndReason() *DocumentAdjustment1 {
	newValue := new(DocumentAdjustment1)
	r.AdjustmentAmountAndReason = append(r.AdjustmentAmountAndReason, newValue)
	return newValue
}

func (r *RemittanceAmount2) SetRemittedAmount(value, currency string) {
	r.RemittedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
