package model

// Information about discounts or waivers to charges and commissions.
type ChargeOrCommissionDiscount1 struct {

	// Difference between the standard fee (charge/commission) amount and the applied fee (charge/commission) amount.
	// EXAMPLE:
	// Standard charge is EUR 100
	// Discount is EUR 30
	// Applied charge is EUR 70.
	Amount *ActiveCurrencyAndAmount `xml:"Amt,omitempty"`

	// Difference between the standard fee (charge/commission) rate and the applied rate of the fee (charge/commission).
	// EXAMPLE:
	// Standard rate is 5%
	// Discount rate is 3%
	// Applied rate is 2%
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Form of the discount or rebate, for example, cash.
	Basis *WaivingInstruction2Choice `xml:"Bsis,omitempty"`
}

func (c *ChargeOrCommissionDiscount1) SetAmount(value, currency string) {
	c.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *ChargeOrCommissionDiscount1) SetRate(value string) {
	c.Rate = (*PercentageRate)(&value)
}

func (c *ChargeOrCommissionDiscount1) AddBasis() *WaivingInstruction2Choice {
	c.Basis = new(WaivingInstruction2Choice)
	return c.Basis
}
