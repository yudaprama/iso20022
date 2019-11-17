package model

// Amount of money associated with a service.
type Fee2 struct {

	// Type of fee (charge/commission).
	Type *ChargeType5Choice `xml:"Tp"`

	// Method used to calculate the fee (charge/commission).
	Basis *ChargeBasis2Choice `xml:"Bsis,omitempty"`

	// Standard fee (charge/commission) amount as specified in the fund prospectus or agreed for the account.
	StandardAmount *ActiveCurrencyAndAmount `xml:"StdAmt,omitempty"`

	// Standard fee (charge/commission) rate used to calculate the amount of the charge or fee, as specified in the fund prospectus or agreed for the account.
	StandardRate *PercentageRate `xml:"StdRate,omitempty"`

	// Discount or waiver applied to the fee (charge/commission).
	DiscountDetails *ChargeOrCommissionDiscount1 `xml:"DscntDtls,omitempty"`

	// Fee (charge/commission) amount applied to the transaction.
	AppliedAmount *ActiveCurrencyAndAmount `xml:"ApldAmt,omitempty"`

	// Final rate used to calculate the fee (charge/commission) amount.
	AppliedRate *PercentageRate `xml:"ApldRate,omitempty"`

	// Reference to a sales agreement that overrides normal processing or the Service Level Agreement (SLA), such as a fee (charge/commission).
	NonStandardSLAReference *Max35Text `xml:"NonStdSLARef,omitempty"`

	// Party entitled to the amount of money resulting from a fee (charge/commission).
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`

	// Indicates the information is provided for information purposes only. When the value is ‘false’ or ‘0’ the amount provided is taken into consideration in the transaction overhead. When the value is ‘true’ or ‘1’ the amount provided is not taken into consideration in the transaction overhead.
	InformativeIndicator *YesNoIndicator `xml:"InftvInd"`
}

func (f *Fee2) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee2) AddBasis() *ChargeBasis2Choice {
	f.Basis = new(ChargeBasis2Choice)
	return f.Basis
}

func (f *Fee2) SetStandardAmount(value, currency string) {
	f.StandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee2) SetStandardRate(value string) {
	f.StandardRate = (*PercentageRate)(&value)
}

func (f *Fee2) AddDiscountDetails() *ChargeOrCommissionDiscount1 {
	f.DiscountDetails = new(ChargeOrCommissionDiscount1)
	return f.DiscountDetails
}

func (f *Fee2) SetAppliedAmount(value, currency string) {
	f.AppliedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee2) SetAppliedRate(value string) {
	f.AppliedRate = (*PercentageRate)(&value)
}

func (f *Fee2) SetNonStandardSLAReference(value string) {
	f.NonStandardSLAReference = (*Max35Text)(&value)
}

func (f *Fee2) AddRecipientIdentification() *PartyIdentification113 {
	f.RecipientIdentification = new(PartyIdentification113)
	return f.RecipientIdentification
}

func (f *Fee2) SetInformativeIndicator(value string) {
	f.InformativeIndicator = (*YesNoIndicator)(&value)
}
