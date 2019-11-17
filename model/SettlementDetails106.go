package model

// Details of settlement of a transaction.
type SettlementDetails106 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition22Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails106) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails106) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails106) AddSettlementTransactionCondition() *SettlementTransactionCondition22Choice {
	newValue := new(SettlementTransactionCondition22Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails106) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails106) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails106) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails106) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails106) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails106) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails106) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails106) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails106) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails106) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails106) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails106) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails106) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails106) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails106) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails106) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails106) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}
