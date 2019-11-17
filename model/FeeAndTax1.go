package model

// Amount of money associated with a service.
type FeeAndTax1 struct {

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Individual fee (charge/commission).
	IndividualFee []*Fee1 `xml:"IndvFee,omitempty"`

	// Individual tax amount.
	IndividualTax []*Tax30 `xml:"IndvTax,omitempty"`
}

func (f *FeeAndTax1) SetCommercialAgreementReference(value string) {
	f.CommercialAgreementReference = (*Max35Text)(&value)
}

func (f *FeeAndTax1) AddIndividualFee() *Fee1 {
	newValue := new(Fee1)
	f.IndividualFee = append(f.IndividualFee, newValue)
	return newValue
}

func (f *FeeAndTax1) AddIndividualTax() *Tax30 {
	newValue := new(Tax30)
	f.IndividualTax = append(f.IndividualTax, newValue)
	return newValue
}
