package model

// Details of settlement of a transaction.
type SettlementDetails137 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType34Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition29Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn4Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction5Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType26Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking5Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed5Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails137) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails137) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails137) AddSecuritiesTransactionType() *SecuritiesTransactionType34Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType34Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails137) AddSettlementTransactionCondition() *SettlementTransactionCondition29Choice {
	newValue := new(SettlementTransactionCondition29Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails137) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails137) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails137) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails137) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails137) AddDeliveryReturnReason() *DeliveryReturn4Choice {
	s.DeliveryReturnReason = new(DeliveryReturn4Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails137) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails137) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails137) AddFXStandingInstruction() *FXStandingInstruction5Choice {
	s.FXStandingInstruction = new(FXStandingInstruction5Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails137) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails137) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails137) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails137) AddRepurchaseType() *RepurchaseType26Choice {
	s.RepurchaseType = new(RepurchaseType26Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails137) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails137) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails137) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails137) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails137) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails137) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails137) AddTracking() *Tracking5Choice {
	s.Tracking = new(Tracking5Choice)
	return s.Tracking
}

func (s *SettlementDetails137) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails137) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails137) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails137) AddModificationCancellationAllowed() *ModificationCancellationAllowed5Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed5Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails137) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails137) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails137) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}
