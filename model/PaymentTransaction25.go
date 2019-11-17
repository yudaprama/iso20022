package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction25 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut5Choice `xml:"CshInOrOut,omitempty"`
}

func (p *PaymentTransaction25) AddCashInOrOut() *CashInOrOut5Choice {
	p.CashInOrOut = new(CashInOrOut5Choice)
	return p.CashInOrOut
}
