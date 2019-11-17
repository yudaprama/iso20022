package model

// Provides information about cash entitlements.
type CashEntitlement1 struct {

	// Entitled cash amount.
	CashAmount *ActiveCurrencyAndAmount `xml:"CshAmt"`
}

func (c *CashEntitlement1) SetCashAmount(value, currency string) {
	c.CashAmount = NewActiveCurrencyAndAmount(value, currency)
}
