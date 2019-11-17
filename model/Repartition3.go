package model

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition3 struct {

	// Amount, units or percentage of financial instrument invested or withdrawn.
	Quantity *UnitsOrAmountOrPercentage1Choice `xml:"Qty"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument29 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *CurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition3) AddQuantity() *UnitsOrAmountOrPercentage1Choice {
	r.Quantity = new(UnitsOrAmountOrPercentage1Choice)
	return r.Quantity
}

func (r *Repartition3) AddFinancialInstrument() *FinancialInstrument29 {
	r.FinancialInstrument = new(FinancialInstrument29)
	return r.FinancialInstrument
}

func (r *Repartition3) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*CurrencyCode)(&value)
}
