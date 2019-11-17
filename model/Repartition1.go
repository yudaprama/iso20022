package model

// Indicates how the amount of the investment plan is split amongst the funds.
type Repartition1 struct {

	// Percentage of amount invested in a funds.
	Percentage *PercentageRate `xml:"Pctg"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrument *FinancialInstrument10 `xml:"FinInstrm"`
}

func (r *Repartition1) SetPercentage(value string) {
	r.Percentage = (*PercentageRate)(&value)
}

func (r *Repartition1) AddFinancialInstrument() *FinancialInstrument10 {
	r.FinancialInstrument = new(FinancialInstrument10)
	return r.FinancialInstrument
}
