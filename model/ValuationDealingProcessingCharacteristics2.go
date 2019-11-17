package model

// Valuation dealing processing characteristics linked to the instrument, ie, not to  the market.
type ValuationDealingProcessingCharacteristics2 struct {

	// Frequency of the valuation.
	ValuationFrequency *EventFrequency5Code `xml:"ValtnFrqcy"`

	// Further details regarding the dealing frequency, eg, Tuesday (for weekly dealing) or last business day of the month.
	ValuationFrequencyDescription *Max350Text `xml:"ValtnFrqcyDesc"`

	// Number of decimal places to which quantities of units/shares are rounded.
	DecimalisationUnits *Number `xml:"DcmlstnUnits"`

	// Number of decimal places to which quantities of units/shares are rounded.
	DecimalisationPrice *Number `xml:"DcmlstnPric"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd"`

	// Type of pricing calculation method.
	PriceMethod *PriceMethod1Code `xml:"PricMtd"`

	// Currencies in which the prices for the investment fund class are published by the fund management company.
	PriceCurrency []*ActiveCurrencyCode `xml:"PricCcy"`
}

func (v *ValuationDealingProcessingCharacteristics2) SetValuationFrequency(value string) {
	v.ValuationFrequency = (*EventFrequency5Code)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) SetValuationFrequencyDescription(value string) {
	v.ValuationFrequencyDescription = (*Max350Text)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) SetDecimalisationUnits(value string) {
	v.DecimalisationUnits = (*Number)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) SetDecimalisationPrice(value string) {
	v.DecimalisationPrice = (*Number)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) SetDualFundIndicator(value string) {
	v.DualFundIndicator = (*YesNoIndicator)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) SetPriceMethod(value string) {
	v.PriceMethod = (*PriceMethod1Code)(&value)
}

func (v *ValuationDealingProcessingCharacteristics2) AddPriceCurrency(value string) {
	v.PriceCurrency = append(v.PriceCurrency, (*ActiveCurrencyCode)(&value))
}
