package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction17 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer, direct debit, investment account or payment card.
	PaymentInstrument *PaymentInstrument9Choice `xml:"PmtInstrm,omitempty"`
}

func (p *PaymentTransaction17) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction17) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction17) AddPaymentInstrument() *PaymentInstrument9Choice {
	p.PaymentInstrument = new(PaymentInstrument9Choice)
	return p.PaymentInstrument
}
