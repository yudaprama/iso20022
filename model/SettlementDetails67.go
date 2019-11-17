package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails67 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator4 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric1Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType9Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition12Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade1Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility1Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn1Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType10Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction1Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility1Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType1Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction1Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity4Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod1Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking1Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing1Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed1Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification20 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification20 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails67) AddHoldIndicator() *HoldIndicator4 {
	s.HoldIndicator = new(HoldIndicator4)
	return s.HoldIndicator
}

func (s *SettlementDetails67) AddPriority() *PriorityNumeric1Choice {
	s.Priority = new(PriorityNumeric1Choice)
	return s.Priority
}

func (s *SettlementDetails67) AddSecuritiesTransactionType() *SecuritiesTransactionType9Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType9Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails67) AddSettlementTransactionCondition() *SettlementTransactionCondition12Choice {
	newValue := new(SettlementTransactionCondition12Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails67) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails67) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails67) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails67) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails67) AddDeliveryReturnReason() *DeliveryReturn1Choice {
	s.DeliveryReturnReason = new(DeliveryReturn1Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails67) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails67) AddExposureType() *ExposureType10Choice {
	s.ExposureType = new(ExposureType10Choice)
	return s.ExposureType
}

func (s *SettlementDetails67) AddFXStandingInstruction() *FXStandingInstruction1Choice {
	s.FXStandingInstruction = new(FXStandingInstruction1Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails67) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails67) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails67) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails67) AddRepurchaseType() *RepurchaseType1Choice {
	s.RepurchaseType = new(RepurchaseType1Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails67) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails67) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails67) AddSettlingCapacity() *SettlingCapacity4Choice {
	s.SettlingCapacity = new(SettlingCapacity4Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails67) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails67) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails67) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails67) AddTracking() *Tracking1Choice {
	s.Tracking = new(Tracking1Choice)
	return s.Tracking
}

func (s *SettlementDetails67) AddAutomaticBorrowing() *AutomaticBorrowing1Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing1Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails67) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails67) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails67) AddModificationCancellationAllowed() *ModificationCancellationAllowed1Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed1Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails67) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails67) AddSecuritiesSubBalanceType() *GenericIdentification20 {
	s.SecuritiesSubBalanceType = new(GenericIdentification20)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails67) AddCashSubBalanceType() *GenericIdentification20 {
	s.CashSubBalanceType = new(GenericIdentification20)
	return s.CashSubBalanceType
}
