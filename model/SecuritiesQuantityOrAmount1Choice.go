package model

// Choice between securities quantities or an amount.
type SecuritiesQuantityOrAmount1Choice struct {

	// Provides information about securities quantity linked to a corporate action option.
	SecuritiesQuantity *SecuritiesOption54 `xml:"SctiesQty"`

	// Cash amount to be instructed.
	InstructedAmount *RestrictedFINActiveCurrencyAndAmount `xml:"InstdAmt"`
}

func (s *SecuritiesQuantityOrAmount1Choice) AddSecuritiesQuantity() *SecuritiesOption54 {
	s.SecuritiesQuantity = new(SecuritiesOption54)
	return s.SecuritiesQuantity
}

func (s *SecuritiesQuantityOrAmount1Choice) SetInstructedAmount(value, currency string) {
	s.InstructedAmount = NewRestrictedFINActiveCurrencyAndAmount(value, currency)
}
