package model

// Provides details about the collateral held, in transit or that still needs to be agreed by both parties.
type MarginCollateral1 struct {

	// Post haircut market value of all margin collateral held by party A.
	HeldByPartyA *ActiveCurrencyAndAmount `xml:"HeldByPtyA,omitempty"`

	// Post haircut market value of all margin collateral held by party B.
	HeldByPartyB *ActiveCurrencyAndAmount `xml:"HeldByPtyB,omitempty"`

	// Sum of all margin agreed amounts due to party A from prior days’ collateral calls where collateral movements have not yet been agreed.
	PriorAgreedToPartyA *ActiveCurrencyAndAmount `xml:"PrrAgrdToPtyA,omitempty"`

	// Sum of all margin agreed amounts due to party B from prior days’ collateral calls where collateral movements have not yet been agreed.
	PriorAgreedToPartyB *ActiveCurrencyAndAmount `xml:"PrrAgrdToPtyB,omitempty"`

	// Sum of all margin collateral movements due to party A in progress but not yet settled.
	InTransitToPartyA *ActiveCurrencyAndAmount `xml:"InTrnstToPtyA,omitempty"`

	// Sum of all margin collateral movements due to party B in progress but not yet settled.
	InTransitToPartyB *ActiveCurrencyAndAmount `xml:"InTrnstToPtyB,omitempty"`
}

func (m *MarginCollateral1) SetHeldByPartyA(value, currency string) {
	m.HeldByPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCollateral1) SetHeldByPartyB(value, currency string) {
	m.HeldByPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCollateral1) SetPriorAgreedToPartyA(value, currency string) {
	m.PriorAgreedToPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCollateral1) SetPriorAgreedToPartyB(value, currency string) {
	m.PriorAgreedToPartyB = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCollateral1) SetInTransitToPartyA(value, currency string) {
	m.InTransitToPartyA = NewActiveCurrencyAndAmount(value, currency)
}

func (m *MarginCollateral1) SetInTransitToPartyB(value, currency string) {
	m.InTransitToPartyB = NewActiveCurrencyAndAmount(value, currency)
}
