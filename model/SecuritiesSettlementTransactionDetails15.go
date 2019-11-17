package model

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails15 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters6 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails32 `xml:"TradDtls,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount29 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails68 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties26 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties26 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties17 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection37 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters6 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters6)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails15) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails15) AddTradeDetails() *SecuritiesTradeDetails32 {
	s.TradeDetails = new(SecuritiesTradeDetails32)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails15) AddQuantityAndAccountDetails() *QuantityAndAccount29 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount29)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementParameters() *SettlementDetails68 {
	s.SettlementParameters = new(SettlementDetails68)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails15) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddDeliveringSettlementParties() *SettlementParties26 {
	s.DeliveringSettlementParties = new(SettlementParties26)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddReceivingSettlementParties() *SettlementParties26 {
	s.ReceivingSettlementParties = new(SettlementParties26)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddCashParties() *CashParties17 {
	s.CashParties = new(CashParties17)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddSettlementAmount() *AmountAndDirection37 {
	s.SettlementAmount = new(AmountAndDirection37)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails15) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails15) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails15) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails15) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
