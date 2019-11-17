package model

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition2 struct {

	// Percentage of amount invested in a funds.
	Percentage *PercentageRate `xml:"Pctg"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument29 `xml:"FinInstrm"`

	// When a fund has multiple currencies within same ISIN, this indicates the currency of the savings or withdrawal plan.
	CurrencyOfPlan *CurrencyCode `xml:"CcyOfPlan,omitempty"`
}

func (r *Repartition2) SetPercentage(value string) {
	r.Percentage = (*PercentageRate)(&value)
}

func (r *Repartition2) AddFinancialInstrument() *FinancialInstrument29 {
	r.FinancialInstrument = new(FinancialInstrument29)
	return r.FinancialInstrument
}

func (r *Repartition2) SetCurrencyOfPlan(value string) {
	r.CurrencyOfPlan = (*CurrencyCode)(&value)
}
