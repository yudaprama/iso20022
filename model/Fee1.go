package model

// Amount of money associated with a service.
type Fee1 struct {

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

	// Requested fee (charge/commission) amount as agreed for the account.
	RequestedAmount *ActiveCurrencyAndAmount `xml:"ReqdAmt,omitempty"`

	// Requested rate used to calculate the amount of the fee (charge/commission), as agreed for the account.
	RequestedRate *PercentageRate `xml:"ReqdRate,omitempty"`

	// Reference to a sales agreement that overrides normal processing or the Service Level Agreement (SLA), such as a fee (charge/commission).
	NonStandardSLAReference *Max35Text `xml:"NonStdSLARef,omitempty"`

	// Party entitled to the amount of money resulting from a fee (charge/commission).
	RecipientIdentification *PartyIdentification113 `xml:"RcptId,omitempty"`
}

func (f *Fee1) AddType() *ChargeType5Choice {
	f.Type = new(ChargeType5Choice)
	return f.Type
}

func (f *Fee1) AddBasis() *ChargeBasis2Choice {
	f.Basis = new(ChargeBasis2Choice)
	return f.Basis
}

func (f *Fee1) SetStandardAmount(value, currency string) {
	f.StandardAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee1) SetStandardRate(value string) {
	f.StandardRate = (*PercentageRate)(&value)
}

func (f *Fee1) AddDiscountDetails() *ChargeOrCommissionDiscount1 {
	f.DiscountDetails = new(ChargeOrCommissionDiscount1)
	return f.DiscountDetails
}

func (f *Fee1) SetRequestedAmount(value, currency string) {
	f.RequestedAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *Fee1) SetRequestedRate(value string) {
	f.RequestedRate = (*PercentageRate)(&value)
}

func (f *Fee1) SetNonStandardSLAReference(value string) {
	f.NonStandardSLAReference = (*Max35Text)(&value)
}

func (f *Fee1) AddRecipientIdentification() *PartyIdentification113 {
	f.RecipientIdentification = new(PartyIdentification113)
	return f.RecipientIdentification
}
