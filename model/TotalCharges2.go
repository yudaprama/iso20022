package model

// Total amount of charges.
type TotalCharges2 struct {

	// Total value of the charges for a specific order.
	TotalAmountOfCharges *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfChrgs,omitempty"`

	// Information related to a specific charge.
	ChargeDetails []*Charge10 `xml:"ChrgDtls"`
}

func (t *TotalCharges2) SetTotalAmountOfCharges(value, currency string) {
	t.TotalAmountOfCharges = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCharges2) AddChargeDetails() *Charge10 {
	newValue := new(Charge10)
	t.ChargeDetails = append(t.ChargeDetails, newValue)
	return newValue
}
