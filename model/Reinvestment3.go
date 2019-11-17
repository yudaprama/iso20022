package model

// Reinvestment information.
type Reinvestment3 struct {

	// Investment fund for the reinvestment.
	FinancialInstrumentDetails *FinancialInstrument56 `xml:"FinInstrmDtls"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Percentage of the reinvestment.
	ReinvestmentPercentage *PercentageRate `xml:"RinvstmtPctg"`
}

func (r *Reinvestment3) AddFinancialInstrumentDetails() *FinancialInstrument56 {
	r.FinancialInstrumentDetails = new(FinancialInstrument56)
	return r.FinancialInstrumentDetails
}

func (r *Reinvestment3) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *Reinvestment3) SetReinvestmentPercentage(value string) {
	r.ReinvestmentPercentage = (*PercentageRate)(&value)
}
