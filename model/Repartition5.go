package model

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition5 struct {

	// Amount, units or percentage of financial instrument invested or withdrawn.
	Quantity *UnitsOrAmountOrPercentage1Choice `xml:"Qty"`

	// Detailed information about the security or investment fund.
	FinancialInstrument *FinancialInstrument56 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *ActiveOrHistoricCurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition5) AddQuantity() *UnitsOrAmountOrPercentage1Choice {
	r.Quantity = new(UnitsOrAmountOrPercentage1Choice)
	return r.Quantity
}

func (r *Repartition5) AddFinancialInstrument() *FinancialInstrument56 {
	r.FinancialInstrument = new(FinancialInstrument56)
	return r.FinancialInstrument
}

func (r *Repartition5) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*ActiveOrHistoricCurrencyCode)(&value)
}
