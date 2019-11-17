package model

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails22 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters14 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages38 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails52 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes64 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount43 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails94 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction11 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties36 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties36 `xml:"RcvgSttlmPties,omitempty"`

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

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters14 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters14)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails22) AddLinkages() *Linkages38 {
	newValue := new(Linkages38)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails22) AddTradeDetails() *SecuritiesTradeDetails52 {
	s.TradeDetails = new(SecuritiesTradeDetails52)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails22) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes64 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes64)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails22) AddQuantityAndAccountDetails() *QuantityAndAccount43 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount43)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementParameters() *SettlementDetails94 {
	s.SettlementParameters = new(SettlementDetails94)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails22) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction11 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction11)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddDeliveringSettlementParties() *SettlementParties36 {
	s.DeliveringSettlementParties = new(SettlementParties36)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddReceivingSettlementParties() *SettlementParties36 {
	s.ReceivingSettlementParties = new(SettlementParties36)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddCashParties() *CashParties26 {
	s.CashParties = new(CashParties26)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddSettlementAmount() *AmountAndDirection45 {
	s.SettlementAmount = new(AmountAndDirection45)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails22) AddOtherAmounts() *OtherAmounts28 {
	s.OtherAmounts = new(OtherAmounts28)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails22) AddOtherBusinessParties() *OtherParties27 {
	s.OtherBusinessParties = new(OtherParties27)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails22) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters4 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters4)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails22) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
