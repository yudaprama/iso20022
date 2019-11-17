package model

// Provides details on the margin result.
type MarginResult1Choice struct {

	// Excess amount that the central counterparty may restitute to the Clearing member.
	ExcessAmount *ActiveCurrencyAndAmount `xml:"XcssAmt"`

	// Deficit amount that the central counterparty will provide to the clearing member.
	DeficitAmount *ActiveCurrencyAndAmount `xml:"DfcitAmt"`
}

func (m *MarginResult1Choice) SetExcessAmount(value, currency string) {
	m.ExcessAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginResult1Choice) SetDeficitAmount(value, currency string) {
	m.DeficitAmount = NewActiveCurrencyAndAmount(value, currency)
}
