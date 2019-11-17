package model

// Choice of payment instruments.
type PaymentInstrument18Choice struct {

	// Settlement instructions for a payment by card.
	PaymentCardDetails *PaymentCard18 `xml:"PmtCardDtls"`

	// Settlement instructions for a payment by direct debit.
	DirectDebitDetails *DirectDebitMandate5 `xml:"DrctDbtDtls"`

	// Indicates whether the payment is done via cheque.
	Cheque *YesNoIndicator `xml:"Chq"`

	// Indicates whether the payment is done via draft.
	BankersDraft *YesNoIndicator `xml:"BkrsDrft"`
}

func (p *PaymentInstrument18Choice) AddPaymentCardDetails() *PaymentCard18 {
	p.PaymentCardDetails = new(PaymentCard18)
	return p.PaymentCardDetails
}

func (p *PaymentInstrument18Choice) AddDirectDebitDetails() *DirectDebitMandate5 {
	p.DirectDebitDetails = new(DirectDebitMandate5)
	return p.DirectDebitDetails
}

func (p *PaymentInstrument18Choice) SetCheque(value string) {
	p.Cheque = (*YesNoIndicator)(&value)
}

func (p *PaymentInstrument18Choice) SetBankersDraft(value string) {
	p.BankersDraft = (*YesNoIndicator)(&value)
}
