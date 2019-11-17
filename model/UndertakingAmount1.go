package model

// Information about an amount.
type UndertakingAmount1 struct {

	// Amount and currency of the undertaking.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Percentage by which the amount claimed under the undertaking may be more than the undertaking amount.
	PlusTolerance *PercentageRate `xml:"PlusTlrnce,omitempty"`

	// Additional information concerning the undertaking amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount1) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount1) SetPlusTolerance(value string) {
	u.PlusTolerance = (*PercentageRate)(&value)
}

func (u *UndertakingAmount1) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
