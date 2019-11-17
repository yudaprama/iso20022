package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction18 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between types of payment instrument, ie, cheque, credit transfer or investment account.
	PaymentInstrument *PaymentInstrument7Choice `xml:"PmtInstrm"`
}

func (p *PaymentTransaction18) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction18) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction18) AddPaymentInstrument() *PaymentInstrument7Choice {
	p.PaymentInstrument = new(PaymentInstrument7Choice)
	return p.PaymentInstrument
}
