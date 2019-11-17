package model

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument9 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Cash account to credit for the payment of the dividends or of the redeemed investments funds.
	CashAccountDetails []*CashAccount4 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by cheque.
	ChequeDetails *Cheque4 `xml:"ChqDtls"`

	// Settlement instructions for a payment by draft.
	BankersDraftDetails *Cheque4 `xml:"BkrsDrftDtls"`
}

func (p *PaymentInstrument9) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument9) AddCashAccountDetails() *CashAccount4 {
	newValue := new(CashAccount4)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument9) AddChequeDetails() *Cheque4 {
	p.ChequeDetails = new(Cheque4)
	return p.ChequeDetails
}

func (p *PaymentInstrument9) AddBankersDraftDetails() *Cheque4 {
	p.BankersDraftDetails = new(Cheque4)
	return p.BankersDraftDetails
}
