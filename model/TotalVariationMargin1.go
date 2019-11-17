package model

// Margin required to cover the risk because of the price fluctuations occurred on the unsettled exposures towards the central counterparty.
type TotalVariationMargin1 struct {

	// Specifies whether the variation margin position is short or long, that is, wether the balance is a negative or positive balance.
	ShortLongIndicator *ShortLong1Code `xml:"ShrtLngInd,omitempty"`

	// Provides the variation margin amount in the reporting currency and optionally in the original currency.
	AmountDetails *Amount2 `xml:"AmtDtls"`
}

func (t *TotalVariationMargin1) SetShortLongIndicator(value string) {
	t.ShortLongIndicator = (*ShortLong1Code)(&value)
}

func (t *TotalVariationMargin1) AddAmountDetails() *Amount2 {
	t.AmountDetails = new(Amount2)
	return t.AmountDetails
}
