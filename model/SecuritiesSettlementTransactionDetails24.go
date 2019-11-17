package model

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails24 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters17 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages49 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails65 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes78 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount54 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails113 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction12 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties58 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties58 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties30 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection85 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts35 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties29 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters5 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters17 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters17)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails24) AddLinkages() *Linkages49 {
	newValue := new(Linkages49)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails24) AddTradeDetails() *SecuritiesTradeDetails65 {
	s.TradeDetails = new(SecuritiesTradeDetails65)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes78 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes78)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails24) AddQuantityAndAccountDetails() *QuantityAndAccount54 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount54)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementParameters() *SettlementDetails113 {
	s.SettlementParameters = new(SettlementDetails113)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails24) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction12 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction12)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddDeliveringSettlementParties() *SettlementParties58 {
	s.DeliveringSettlementParties = new(SettlementParties58)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddReceivingSettlementParties() *SettlementParties58 {
	s.ReceivingSettlementParties = new(SettlementParties58)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddCashParties() *CashParties30 {
	s.CashParties = new(CashParties30)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddSettlementAmount() *AmountAndDirection85 {
	s.SettlementAmount = new(AmountAndDirection85)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails24) AddOtherAmounts() *OtherAmounts35 {
	s.OtherAmounts = new(OtherAmounts35)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails24) AddOtherBusinessParties() *OtherParties29 {
	s.OtherBusinessParties = new(OtherParties29)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails24) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters5 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters5)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails24) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
