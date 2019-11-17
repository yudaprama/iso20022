package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction19 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, credit transfer, cheque, payment card, investment cash account or direct debit.
	PaymentInstrument *PaymentInstrument10Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction19) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction19) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction19) AddPaymentInstrument() *PaymentInstrument10Choice {
	p.PaymentInstrument = new(PaymentInstrument10Choice)
	return p.PaymentInstrument
}
