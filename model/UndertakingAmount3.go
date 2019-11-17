package model

// Information about an amount.
type UndertakingAmount3 struct {

	// Amount and currency of the demand.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Additional information concerning the demand amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount3) SetAmount(value, currency string) {
	u.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount3) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
