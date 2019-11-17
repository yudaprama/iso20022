package model

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails27 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters13 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages40 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails50 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount44 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails122 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties39 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties39 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties26 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection45 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts28 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties27 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters4 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters13 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters13)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails27) AddLinkages() *Linkages40 {
	newValue := new(Linkages40)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails27) AddTradeDetails() *SecuritiesTradeDetails50 {
	s.TradeDetails = new(SecuritiesTradeDetails50)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails27) AddQuantityAndAccountDetails() *QuantityAndAccount44 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount44)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementParameters() *SettlementDetails122 {
	s.SettlementParameters = new(SettlementDetails122)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails27) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddDeliveringSettlementParties() *SettlementParties39 {
	s.DeliveringSettlementParties = new(SettlementParties39)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddReceivingSettlementParties() *SettlementParties39 {
	s.ReceivingSettlementParties = new(SettlementParties39)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails27) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails27) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails27) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails27) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
