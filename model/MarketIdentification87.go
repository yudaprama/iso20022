package model

// Identifies the market and purpose for the settlement.
type MarketIdentification87 struct {

	// Country in which the financial instrument is to be settled.
	Country *CountryCode `xml:"Ctry"`

	// Type of instrument covered by the SSI instruction.
	ClassificationType *ClassificationType1Choice `xml:"ClssfctnTp"`

	// Purpose of the instruction,  for example, whether for regular payments, margin payments related to a collateral movement, securities settlements , securities lending.
	SettlementPurpose *Purpose3Choice `xml:"SttlmPurp,omitempty"`
}

func (m *MarketIdentification87) SetCountry(value string) {
	m.Country = (*CountryCode)(&value)
}

func (m *MarketIdentification87) AddClassificationType() *ClassificationType1Choice {
	m.ClassificationType = new(ClassificationType1Choice)
	return m.ClassificationType
}

func (m *MarketIdentification87) AddSettlementPurpose() *Purpose3Choice {
	m.SettlementPurpose = new(Purpose3Choice)
	return m.SettlementPurpose
}
