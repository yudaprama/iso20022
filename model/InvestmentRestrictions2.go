package model

// Investment restrictions linked to the instrument.
type InvestmentRestrictions2 struct {

	// Minimum initial quantity of securities, expressed as an amount that must be purchased at subscription.
	MinimumInitialSubscriptionAmount *ActiveCurrencyAndAmount `xml:"MinInitlSbcptAmt,omitempty"`

	// Minimum initial number of units/shares that must be purchased.
	MinimumInitialSubscriptionUnits *Number `xml:"MinInitlSbcptUnits,omitempty"`

	// Minimum quantity of securities, expressed as an amount that must be purchased.
	MinimumSubsequentSubscriptionAmount *ActiveCurrencyAndAmount `xml:"MinSbsqntSbcptAmt,omitempty"`

	// Minimum quantity of securities, expressed as number of units/shares that must be purchased.
	MinimumSubsequentSubscriptionUnits *Number `xml:"MinSbsqntSbcptUnits,omitempty"`

	// Maximum quantity of securities, expressed as an amount that can be sold.
	MaximumRedemptionAmount *ActiveCurrencyAndAmount `xml:"MaxRedAmt,omitempty"`

	// Maximum number of shares/units that may be redeemed on a single dealing day.
	MaximumRedemptionUnits *Number `xml:"MaxRedUnits,omitempty"`

	// Specifies any other restrictions that may limit an investor's ability to redeem.
	OtherRedemptionRestrictions *Max350Text `xml:"OthrRedRstrctns,omitempty"`

	// Minimum value of units that must be maintained to avoid automatic redemption.
	MinimumHoldingAmount *ActiveCurrencyAndAmount `xml:"MinHldgAmt,omitempty"`

	// Minimum number of units that must be maintained to avoid automatic redemption.
	MinimumHoldingUnits *DecimalNumber `xml:"MinHldgUnits,omitempty"`

	// Description of a period, that may be a number of days, weeks or descriptive period during which the units/shares must be held following their issue before redemption will be permitted.
	MinimumHoldingPeriod *Max70Text `xml:"MinHldgPrd,omitempty"`

	// Indicates whether registered investors are able to transfer some or all of their holdings to third parties.
	HoldingTransferable *HoldingTransferable1Code `xml:"HldgTrfbl"`
}

func (i *InvestmentRestrictions2) SetMinimumInitialSubscriptionAmount(value, currency string) {
	i.MinimumInitialSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentRestrictions2) SetMinimumInitialSubscriptionUnits(value string) {
	i.MinimumInitialSubscriptionUnits = (*Number)(&value)
}

func (i *InvestmentRestrictions2) SetMinimumSubsequentSubscriptionAmount(value, currency string) {
	i.MinimumSubsequentSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentRestrictions2) SetMinimumSubsequentSubscriptionUnits(value string) {
	i.MinimumSubsequentSubscriptionUnits = (*Number)(&value)
}

func (i *InvestmentRestrictions2) SetMaximumRedemptionAmount(value, currency string) {
	i.MaximumRedemptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentRestrictions2) SetMaximumRedemptionUnits(value string) {
	i.MaximumRedemptionUnits = (*Number)(&value)
}

func (i *InvestmentRestrictions2) SetOtherRedemptionRestrictions(value string) {
	i.OtherRedemptionRestrictions = (*Max350Text)(&value)
}

func (i *InvestmentRestrictions2) SetMinimumHoldingAmount(value, currency string) {
	i.MinimumHoldingAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InvestmentRestrictions2) SetMinimumHoldingUnits(value string) {
	i.MinimumHoldingUnits = (*DecimalNumber)(&value)
}

func (i *InvestmentRestrictions2) SetMinimumHoldingPeriod(value string) {
	i.MinimumHoldingPeriod = (*Max70Text)(&value)
}

func (i *InvestmentRestrictions2) SetHoldingTransferable(value string) {
	i.HoldingTransferable = (*HoldingTransferable1Code)(&value)
}
