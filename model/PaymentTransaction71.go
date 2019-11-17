package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction71 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut7Choice `xml:"CshInOrOut"`
}

func (p *PaymentTransaction71) AddCashInOrOut() *CashInOrOut7Choice {
	p.CashInOrOut = new(CashInOrOut7Choice)
	return p.CashInOrOut
}
