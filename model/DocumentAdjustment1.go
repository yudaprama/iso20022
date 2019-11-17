package model

// Set of elements used to provide information on the amount and reason of the document adjustment.
type DocumentAdjustment1 struct {

	// Amount of money of the document adjustment.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt"`

	// Specifies whether the adjustment must be substracted or added to the total amount.
	CreditDebitIndicator *CreditDebitCode `xml:"CdtDbtInd,omitempty"`

	// Specifies the reason for the adjustment.
	Reason *Max4Text `xml:"Rsn,omitempty"`

	// Provides further details on the document adjustment.
	AdditionalInformation *Max140Text `xml:"AddtlInf,omitempty"`
}

func (d *DocumentAdjustment1) SetAmount(value, currency string) {
	d.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DocumentAdjustment1) SetCreditDebitIndicator(value string) {
	d.CreditDebitIndicator = (*CreditDebitCode)(&value)
}

func (d *DocumentAdjustment1) SetReason(value string) {
	d.Reason = (*Max4Text)(&value)
}

func (d *DocumentAdjustment1) SetAdditionalInformation(value string) {
	d.AdditionalInformation = (*Max140Text)(&value)
}
