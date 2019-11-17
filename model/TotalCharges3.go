package model

// Total amount of charges.
type TotalCharges3 struct {

	// Total value of the charges for a specific order.
	TotalAmountOfCharges *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfChrgs,omitempty"`

	// Information related to a specific charge.
	ChargeDetails []*Charge18 `xml:"ChrgDtls"`
}

func (t *TotalCharges3) SetTotalAmountOfCharges(value, currency string) {
	t.TotalAmountOfCharges = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalCharges3) AddChargeDetails() *Charge18 {
	newValue := new(Charge18)
	t.ChargeDetails = append(t.ChargeDetails, newValue)
	return newValue
}
