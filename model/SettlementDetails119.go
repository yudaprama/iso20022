package model

// Details of settlement of a transaction.
type SettlementDetails119 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator6 `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType32Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the reason of a delivery return.
	DeliveryReturnReason *DeliveryReturn3Choice `xml:"DlvryRtrRsn,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the foreign exchange standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction4Choice `xml:"FxStgInstr,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType23Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking4Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed4Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails119) AddHoldIndicator() *HoldIndicator6 {
	s.HoldIndicator = new(HoldIndicator6)
	return s.HoldIndicator
}

func (s *SettlementDetails119) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails119) AddSecuritiesTransactionType() *SecuritiesTransactionType32Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType32Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails119) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails119) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails119) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails119) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails119) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails119) AddDeliveryReturnReason() *DeliveryReturn3Choice {
	s.DeliveryReturnReason = new(DeliveryReturn3Choice)
	return s.DeliveryReturnReason
}

func (s *SettlementDetails119) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails119) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails119) AddFXStandingInstruction() *FXStandingInstruction4Choice {
	s.FXStandingInstruction = new(FXStandingInstruction4Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails119) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails119) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails119) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails119) AddRepurchaseType() *RepurchaseType23Choice {
	s.RepurchaseType = new(RepurchaseType23Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails119) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails119) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails119) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails119) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails119) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails119) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails119) AddTracking() *Tracking4Choice {
	s.Tracking = new(Tracking4Choice)
	return s.Tracking
}

func (s *SettlementDetails119) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails119) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails119) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails119) AddModificationCancellationAllowed() *ModificationCancellationAllowed4Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed4Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails119) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails119) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails119) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}
