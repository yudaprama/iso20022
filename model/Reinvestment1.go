package model

// Reinvestment information.
type Reinvestment1 struct {

	// Investment fund for the reinvestment.
	FundDetails *FinancialInstrument29 `xml:"FndDtls"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Percentage of the reinvestment.
	ReinvestmentPercentage *PercentageRate `xml:"RinvstmtPctg"`
}

func (r *Reinvestment1) AddFundDetails() *FinancialInstrument29 {
	r.FundDetails = new(FinancialInstrument29)
	return r.FundDetails
}

func (r *Reinvestment1) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *Reinvestment1) SetReinvestmentPercentage(value string) {
	r.ReinvestmentPercentage = (*PercentageRate)(&value)
}
