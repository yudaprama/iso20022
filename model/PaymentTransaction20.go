package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction20 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut4Choice `xml:"CshInOrOut"`
}

func (p *PaymentTransaction20) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction20) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction20) AddCashInOrOut() *CashInOrOut4Choice {
	p.CashInOrOut = new(CashInOrOut4Choice)
	return p.CashInOrOut
}
