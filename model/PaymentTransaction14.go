package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction14 struct {

	// Amount of money to be transferred between the debtor and creditor before bank transaction charges.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which the first agent expects the cash to be available to the final agent.
	SettlementDate *ISODate `xml:"SttlmDt,omitempty"`

	// Choice between cash-in or cash-out.
	CashInOrOutChoice *CashInOrOut3Choice `xml:"CshInOrOutChc,omitempty"`
}

func (p *PaymentTransaction14) SetSettlementAmount(value, currency string) {
	p.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PaymentTransaction14) SetSettlementDate(value string) {
	p.SettlementDate = (*ISODate)(&value)
}

func (p *PaymentTransaction14) AddCashInOrOutChoice() *CashInOrOut3Choice {
	p.CashInOrOutChoice = new(CashInOrOut3Choice)
	return p.CashInOrOutChoice
}
