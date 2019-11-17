package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails31 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition7Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails31) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails31) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails31) AddSettlementTransactionCondition() *SettlementTransactionCondition7Choice {
	newValue := new(SettlementTransactionCondition7Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails31) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails31) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails31) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails31) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails31) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails31) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails31) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails31) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails31) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails31) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails31) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails31) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails31) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails31) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails31) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails31) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails31) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}
