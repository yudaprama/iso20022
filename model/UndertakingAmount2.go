package model

// Information about an amount.
type UndertakingAmount2 struct {

	// Choice of amounts.
	AmountChoice *Amount1Choice `xml:"AmtChc"`

	// Additional information concerning the amended amount.
	AdditionalInformation []*Max2000Text `xml:"AddtlInf,omitempty"`
}

func (u *UndertakingAmount2) AddAmountChoice() *Amount1Choice {
	u.AmountChoice = new(Amount1Choice)
	return u.AmountChoice
}

func (u *UndertakingAmount2) AddAdditionalInformation(value string) {
	u.AdditionalInformation = append(u.AdditionalInformation, (*Max2000Text)(&value))
}
