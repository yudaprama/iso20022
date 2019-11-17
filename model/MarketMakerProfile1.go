package model

// Market maker profile.
type MarketMakerProfile1 struct {

	// Period of the contract.
	ContractPeriod *DateTimePeriodDetails1 `xml:"CtrctPrd,omitempty"`

	// Indicates whether the market maker is obligated to comply with the requirements of the contract it holds with the  exchange or is exempt from these obligations.
	Compliance *YesNoIndicator `xml:"Cmplc,omitempty"`

	// Percentage of the variation between the maximum accepted minimum and maximum value of an action.
	MaximumSpread *PercentageRate `xml:"MaxSprd,omitempty"`

	// Rate of discount.
	Discount *PercentageRate `xml:"Dscnt,omitempty"`
}

func (m *MarketMakerProfile1) AddContractPeriod() *DateTimePeriodDetails1 {
	m.ContractPeriod = new(DateTimePeriodDetails1)
	return m.ContractPeriod
}

func (m *MarketMakerProfile1) SetCompliance(value string) {
	m.Compliance = (*YesNoIndicator)(&value)
}

func (m *MarketMakerProfile1) SetMaximumSpread(value string) {
	m.MaximumSpread = (*PercentageRate)(&value)
}

func (m *MarketMakerProfile1) SetDiscount(value string) {
	m.Discount = (*PercentageRate)(&value)
}
