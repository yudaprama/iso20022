package model

// Amount of money associated with a service.
type Fee3 struct {

	// Type of fee (charge/commission).
	Type *ChargeType5Choice `xml:"Tp,omitempty"`

	// Modified value of the standard fee (charge/commission) amount applied on the order (the standard fee (charge/commission) amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedStandardAmount *ActiveCurrencyAndAmount `xml:"RprdStdAmt,omitempty"`

	// Modified value of the standard fee (charge/commission) rate applied on the order (the standard fee (charge/commission) rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedStandardRate *PercentageRate `xml:"RprdStdRate,omitempty"`

	// Modified value of the discount amount applied on the order (the discount amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedDiscountAmount *ActiveCurrencyAndAmount `xml:"RprdDscntAmt,omitempty"`

	// Modified value of the discount rate applied on the order (the discount rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedDiscountRate *PercentageRate `xml:"RprdDscntRate,omitempty"`

	// Modified value of the requested fee (charge/commission) amount applied on the order (the requested fee (charge/commission) amount in the original individual order that has been repaired so that the order can be accepted).
	RepairedRequestedAmount *ActiveCurrencyAndAmount `xml:"RprdReqdAmt,omitempty"`

	// Modified value of the requested  fee (charge/commission) rate applied on the order (the requested fee (charge/commission) rate in the original individual order that has been repaired so that the order can be accepted).
	RepairedRequestedRate *PercentageRate `xml:"RprdReqdRate,omitempty"`

	// Reference to the agreement established between the fund and another party. This element, amongst others, defines the conditions of the commissions.
	CommercialAgreementReference *Max35Text `xml:"ComrclAgrmtRef,omitempty"`

	// Indicates if the CommercialAgreementReference is a new reference or not.
	NewCommercialAgreementReferenceIndicator *YesNoIndicator `xml:"NewComrclAgrmtRefInd,omitempty"`
}

func (f *Fee3) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee3) SetRepairedStandardAmount(value, currency string) {
	f.RepairedStandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedStandardRate(value string) {
	f.RepairedStandardRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetRepairedDiscountAmount(value, currency string) {
	f.RepairedDiscountAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedDiscountRate(value string) {
	f.RepairedDiscountRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetRepairedRequestedAmount(value, currency string) {
	f.RepairedRequestedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee3) SetRepairedRequestedRate(value string) {
	f.RepairedRequestedRate = (*PercentageRate)(&value)
}

func (f *Fee3) SetCommercialAgreementReference(value string) {
	f.CommercialAgreementReference = (*Max35Text)(&value)
}

func (f *Fee3) SetNewCommercialAgreementReferenceIndicator(value string) {
	f.NewCommercialAgreementReferenceIndicator = (*YesNoIndicator)(&value)
}
