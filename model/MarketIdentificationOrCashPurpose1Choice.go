package model

// Choice of cash purpose or a securities market identifier.
type MarketIdentificationOrCashPurpose1Choice struct {

	// Identifies the market for the settlement. This consists of the country code and the asset class. For example, if the SSI is for equities in the DTCC, the country code is ‘US’ and the classification type is ‘equity’.
	SettlementInstructionMarketIdentification *MarketIdentification87 `xml:"SttlmInstrMktId"`

	// Underlying reason for the payment SSI instruction, expressed as a code.
	CashSSIPurpose []*ExternalMarketArea1Code `xml:"CshSSIPurp"`
}

func (m *MarketIdentificationOrCashPurpose1Choice) AddSettlementInstructionMarketIdentification() *MarketIdentification87 {
	m.SettlementInstructionMarketIdentification = new(MarketIdentification87)
	return m.SettlementInstructionMarketIdentification
}

func (m *MarketIdentificationOrCashPurpose1Choice) AddCashSSIPurpose(value string) {
	m.CashSSIPurpose = append(m.CashSSIPurpose, (*ExternalMarketArea1Code)(&value))
}
