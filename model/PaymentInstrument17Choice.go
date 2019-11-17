package model

// Choice of payment instruments.
type PaymentInstrument17Choice struct {

	// Cash account to debit for the payment of a subscription or of a savings plan to an investment fund.
	CashAccountDetails []*CashAccount26 `xml:"CshAcctDtls"`

	// Settlement instructions for a payment by card.
	PaymentCardDetails *PaymentCard2 `xml:"PmtCardDtls"`

	// Settlement instructions for a payment by direct debit.
	DirectDebitDetails *DirectDebitMandate4 `xml:"DrctDbtDtls"`

	// Indicates whether the payment is done via cheque.
	Cheque *YesNoIndicator `xml:"Chq"`

	// Indicates whether the payment is done via draft.
	BankersDraft *YesNoIndicator `xml:"BkrsDrft"`
}

func (p *PaymentInstrument17Choice) AddCashAccountDetails() *CashAccount26 {
	newValue := new(CashAccount26)
	p.CashAccountDetails = append(p.CashAccountDetails, newValue)
	return newValue
}

func (p *PaymentInstrument17Choice) AddPaymentCardDetails() *PaymentCard2 {
	p.PaymentCardDetails = new(PaymentCard2)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument17Choice) AddDirectDebitDetails() *DirectDebitMandate4 {
	p.DirectDebitDetails = new(DirectDebitMandate4)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument17Choice) SetCheque(value string) {
	p.Cheque = (*YesNoIndicator)(&value)
}

func (p *PaymentInstrument17Choice) SetBankersDraft(value string) {
	p.BankersDraft = (*YesNoIndicator)(&value)
}
