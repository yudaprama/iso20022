package model

// Amount of money associated with a service.
type TotalFeesAndTaxes40 struct {

	// Total amount of overhead applied to the transaction that impacts the settlement amount.
	TotalOverheadApplied *ActiveCurrencyAndAmount `xml:"TtlOvrhdApld,omitempty"`

	// Total amount of fees (charge/commissions) applied to the transaction that impacts the settlement amount.
	TotalFees *ActiveCurrencyAndAmount `xml:"TtlFees,omitempty"`

	// Total amount of taxes applied to the transaction that impacts the settlement amount.
	TotalTaxes *ActiveCurrencyAndAmount `xml:"TtlTaxs,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Individual fee (charge/commission).
	IndividualFee []*Fee2 `xml:"IndvFee,omitempty"`

	// Individual tax.
	IndividualTax []*Tax31 `xml:"IndvTax,omitempty"`
}

func (t *TotalFeesAndTaxes40) SetTotalOverheadApplied(value, currency string) {
	t.TotalOverheadApplied = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalFeesAndTaxes40) SetTotalFees(value, currency string) {
	t.TotalFees = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalFeesAndTaxes40) SetTotalTaxes(value, currency string) {
	t.TotalTaxes = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalFeesAndTaxes40) SetCommercialAgreementReference(value string) {
	t.CommercialAgreementReference = (*Max35Text)(&value)
}

func (t *TotalFeesAndTaxes40) AddIndividualFee() *Fee2 {
	newValue := new(Fee2)
	t.IndividualFee = append(t.IndividualFee, newValue)
	return newValue
}

func (t *TotalFeesAndTaxes40) AddIndividualTax() *Tax31 {
	newValue := new(Tax31)
	t.IndividualTax = append(t.IndividualTax, newValue)
	return newValue
}
