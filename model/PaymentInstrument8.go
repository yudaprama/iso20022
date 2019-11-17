package model

// Instrument that has or represents monetary value and is used to process a payment instruction.
type PaymentInstrument8 struct {

	// Currency associated with the payment instrument.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy"`

	// Cash account to debit for the payment of a subscription or of a savings plan to an investment fund.
	CashAccountDetails []*CashAccount4 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by card.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Settlement instructions for a payment by direct debit.
	DirectDebitDetails *DirectDebitMandate4 `xml:"DrctDbtDtls"`

	// Indicates whether the payment is done via cheque.
	Cheque *YesNoIndicator `xml:"Chq"`

	// Indicates whether the payment is done via draft.
	BankersDraft *YesNoIndicator `xml:"BkrsDrft"`
}

func (p *PaymentInstrument8) SetSettlementCurrency(value string) {
	p.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}

func (p *PaymentInstrument8) AddCashAccountDetails() *CashAccount4 {
	newValue := new(CashAccount4)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument8) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument8) AddDirectDebitDetails() *DirectDebitMandate4 {
	p.DirectDebitDetails = new(DirectDebitMandate4)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument8) SetCheque(value string) {
	p.Cheque = (*YesNoIndicator)(&value)
}

func (p *PaymentInstrument8) SetBankersDraft(value string) {
	p.BankersDraft = (*YesNoIndicator)(&value)
}
