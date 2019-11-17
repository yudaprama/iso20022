package model

// Payment processes required to transfer cash from the debtor to the creditor.
type PaymentTransaction26 struct {

	// Choice between cash-in or cash-out.
	CashInOrOut *CashInOrOut6Choice `xml:"CshInOrOut,omitempty"`
}

func (p *PaymentTransaction26) AddCashInOrOut() *CashInOrOut6Choice {
	p.CashInOrOut = new(CashInOrOut6Choice)
	return p.CashInOrOut
}
