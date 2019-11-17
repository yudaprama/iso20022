package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails71 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails71) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails71) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails71) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails71) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails71) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails71) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails71) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails71) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails71) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails71) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails71) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails71) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails71) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails71) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails71) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails71) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails71) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}
