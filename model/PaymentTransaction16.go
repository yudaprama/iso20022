package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction16 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument8Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction16) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction16) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction16) AddPaymentInstrument() *PaymentInstrument8Choice {
	p.PaymentInstrument = new(PaymentInstrument8Choice)
	return p.PaymentInstrument
}
