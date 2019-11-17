package model

// Provides the details of the update(s) for the settlement transaction.
type SecuritiesSettlementTransactionDetails2 struct {

	// Provides settlement type and identification information.
	SettlementTypeAndAdditionalParameters *SettlementTypeAndAdditionalParameters5 `xml:"SttlmTpAndAddtlParams,omitempty"`

	// Link to another transaction that must be processed after, before or at the same time.
	Linkages []*Linkages1 `xml:"Lnkgs,omitempty"`

	// Details of the trade.
	TradeDetails *SecuritiesTradeDetails26 `xml:"TradDtls,omitempty"`

	// Financial instrument representing a sum of rights of the investor vis-a-vis the issuer.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Attributes defining a financial instrument.
	FinancialInstrumentAttributes *FinancialInstrumentAttributes35 `xml:"FinInstrmAttrbts,omitempty"`

	// Details related to the account and quantity involved in the transaction.
	QuantityAndAccountDetails *QuantityAndAccount30 `xml:"QtyAndAcctDtls,omitempty"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails50 `xml:"SttlmParams,omitempty"`

	// Specifies what settlement standing instruction database is to be used to derive the settlement parties involved in the transaction.
	StandingSettlementInstructionDetails *StandingSettlementInstruction4 `xml:"StgSttlmInstrDtls,omitempty"`

	// Identifies the chain of delivering settlement parties.
	DeliveringSettlementParties *SettlementParties11 `xml:"DlvrgSttlmPties,omitempty"`

	// Identifies the chain of receiving settlement parties.
	ReceivingSettlementParties *SettlementParties11 `xml:"RcvgSttlmPties,omitempty"`

	// Cash parties involved in the transaction if different for the securities settlement parties.
	CashParties *CashParties8 `xml:"CshPties,omitempty"`

	// Total amount of money to be paid or received in exchange for the securities.
	SettlementAmount *AmountAndDirection32 `xml:"SttlmAmt,omitempty"`

	// Other amounts than the settlement amount.
	OtherAmounts *OtherAmounts14 `xml:"OthrAmts,omitempty"`

	// Other business parties relevant to the transaction.
	OtherBusinessParties *OtherParties19 `xml:"OthrBizPties,omitempty"`

	// Provides information required for the registration or physical settlement.
	AdditionalPhysicalOrRegistrationDetails *RegistrationParameters1 `xml:"AddtlPhysOrRegnDtls,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementTypeAndAdditionalParameters() *SettlementTypeAndAdditionalParameters5 {
	s.SettlementTypeAndAdditionalParameters = new(SettlementTypeAndAdditionalParameters5)
	return s.SettlementTypeAndAdditionalParameters
}

func (s *SecuritiesSettlementTransactionDetails2) AddLinkages() *Linkages1 {
	newValue := new(Linkages1)
	s.Linkages = append(s.Linkages, newValue)
	return newValue
}

func (s *SecuritiesSettlementTransactionDetails2) AddTradeDetails() *SecuritiesTradeDetails26 {
	s.TradeDetails = new(SecuritiesTradeDetails26)
	return s.TradeDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	s.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return s.FinancialInstrumentIdentification
}

func (s *SecuritiesSettlementTransactionDetails2) AddFinancialInstrumentAttributes() *FinancialInstrumentAttributes35 {
	s.FinancialInstrumentAttributes = new(FinancialInstrumentAttributes35)
	return s.FinancialInstrumentAttributes
}

func (s *SecuritiesSettlementTransactionDetails2) AddQuantityAndAccountDetails() *QuantityAndAccount30 {
	s.QuantityAndAccountDetails = new(QuantityAndAccount30)
	return s.QuantityAndAccountDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementParameters() *SettlementDetails50 {
	s.SettlementParameters = new(SettlementDetails50)
	return s.SettlementParameters
}

func (s *SecuritiesSettlementTransactionDetails2) AddStandingSettlementInstructionDetails() *StandingSettlementInstruction4 {
	s.StandingSettlementInstructionDetails = new(StandingSettlementInstruction4)
	return s.StandingSettlementInstructionDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddDeliveringSettlementParties() *SettlementParties11 {
	s.DeliveringSettlementParties = new(SettlementParties11)
	return s.DeliveringSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddReceivingSettlementParties() *SettlementParties11 {
	s.ReceivingSettlementParties = new(SettlementParties11)
	return s.ReceivingSettlementParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddCashParties() *CashParties8 {
	s.CashParties = new(CashParties8)
	return s.CashParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddSettlementAmount() *AmountAndDirection32 {
	s.SettlementAmount = new(AmountAndDirection32)
	return s.SettlementAmount
}

func (s *SecuritiesSettlementTransactionDetails2) AddOtherAmounts() *OtherAmounts14 {
	s.OtherAmounts = new(OtherAmounts14)
	return s.OtherAmounts
}

func (s *SecuritiesSettlementTransactionDetails2) AddOtherBusinessParties() *OtherParties19 {
	s.OtherBusinessParties = new(OtherParties19)
	return s.OtherBusinessParties
}

func (s *SecuritiesSettlementTransactionDetails2) AddAdditionalPhysicalOrRegistrationDetails() *RegistrationParameters1 {
	s.AdditionalPhysicalOrRegistrationDetails = new(RegistrationParameters1)
	return s.AdditionalPhysicalOrRegistrationDetails
}

func (s *SecuritiesSettlementTransactionDetails2) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	s.SupplementaryData = append(s.SupplementaryData, newValue)
	return newValue
}
