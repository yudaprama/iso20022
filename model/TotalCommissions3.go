package model

// Total amount of commissions related to a specific order.
type TotalCommissions3 struct {

	// Total value of the commissions for a specific order.
	TotalAmountOfCommissions *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfComssns,omitempty"`

	// Information related to a specific commission.
	CommissionDetails []*Commission10 `xml:"ComssnDtls"`
}

func (t *TotalCommissions3) SetTotalAmountOfCommissions(value, currency string) {
	t.TotalAmountOfCommissions = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCommissions3) AddCommissionDetails() *Commission10 {
	newValue := new(Commission10)
	t.CommissionDetails = append(t.CommissionDetails, newValue)
	return newValue
}
