package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails43 struct {

	// Indicates whether the trade is to be settled as principal or netted off against another trade.
	SettlementTransactionType *SettlementTransactionType1Choice `xml:"SttlmTxTp"`

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Specifies whether the transaction is to be executed with a high priority.
	Priority *PriorityNumeric3Choice `xml:"Prty,omitempty"`

	// Specifies if the Electronic Trade Confirmation (ETC) service provider is to generate or not a settlement instruction.
	SettlementInstructionGeneration *SettlementInstructionGeneration1Choice `xml:"SttlmInstrGnrtn,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition11Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *YesNoIndicator `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership3Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade3Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility3Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, eg, cheque clearing.
	CashClearingSystem *CashSettlementSystem3Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade causing the collateral movement.
	ExposureType *ExposureType9Choice `xml:"XpsrTp,omitempty"`

	// Specifies whether the forex standing instruction in place should apply.
	FXStandingInstruction *FXStandingInstruction3Choice `xml:"FxStgInstr,omitempty"`

	// Account servicer is instructed to buy the indicated currency after the receipt of cash proceeds or to sell the indicated currency in order to obtain the necessary currency to fund the transaction.
	CurrencyToBuyOrSell *CurrencyToBuyOrSell1Choice `xml:"CcyToBuyOrSell,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide3Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility3Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration6Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType11Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction3Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS3Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity3Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod3Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty3Choice `xml:"TaxCpcty,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDutyIndicator *YesNoIndicator `xml:"StmpDtyInd,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification38 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the loan and/or collateral is tracked.
	Tracking *Tracking3Choice `xml:"Trckg,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing5Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee3Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether, for a securities lending/borrowing settlement transaction, the lender will instruct the return leg as agreed with the borrower.
	ReturnLeg *YesNoIndicator `xml:"RtrLeg,omitempty"`

	// Specifies whether a third party is allowed to modify or cancel the transaction.
	ModificationCancellationAllowed *ModificationCancellationAllowed3Choice `xml:"ModCxlAllwd,omitempty"`

	// Specifies whether securities should be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`
}

func (s *SettlementDetails43) AddSettlementTransactionType() *SettlementTransactionType1Choice {
	s.SettlementTransactionType = new(SettlementTransactionType1Choice)
	return s.SettlementTransactionType
}

func (s *SettlementDetails43) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddPriority() *PriorityNumeric3Choice {
	s.Priority = new(PriorityNumeric3Choice)
	return s.Priority
}

func (s *SettlementDetails43) AddSettlementInstructionGeneration() *SettlementInstructionGeneration1Choice {
	s.SettlementInstructionGeneration = new(SettlementInstructionGeneration1Choice)
	return s.SettlementInstructionGeneration
}

func (s *SettlementDetails43) AddSettlementTransactionCondition() *SettlementTransactionCondition11Choice {
	newValue := new(SettlementTransactionCondition11Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails43) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddBeneficialOwnership() *BeneficialOwnership3Choice {
	s.BeneficialOwnership = new(BeneficialOwnership3Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails43) AddBlockTrade() *BlockTrade3Choice {
	s.BlockTrade = new(BlockTrade3Choice)
	return s.BlockTrade
}

func (s *SettlementDetails43) AddCCPEligibility() *CentralCounterPartyEligibility3Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility3Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails43) AddCashClearingSystem() *CashSettlementSystem3Choice {
	s.CashClearingSystem = new(CashSettlementSystem3Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails43) AddExposureType() *ExposureType9Choice {
	s.ExposureType = new(ExposureType9Choice)
	return s.ExposureType
}

func (s *SettlementDetails43) AddFXStandingInstruction() *FXStandingInstruction3Choice {
	s.FXStandingInstruction = new(FXStandingInstruction3Choice)
	return s.FXStandingInstruction
}

func (s *SettlementDetails43) AddCurrencyToBuyOrSell() *CurrencyToBuyOrSell1Choice {
	s.CurrencyToBuyOrSell = new(CurrencyToBuyOrSell1Choice)
	return s.CurrencyToBuyOrSell
}

func (s *SettlementDetails43) AddMarketClientSide() *MarketClientSide3Choice {
	s.MarketClientSide = new(MarketClientSide3Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails43) AddNettingEligibility() *NettingEligibility3Choice {
	s.NettingEligibility = new(NettingEligibility3Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails43) AddRegistration() *Registration6Choice {
	s.Registration = new(Registration6Choice)
	return s.Registration
}

func (s *SettlementDetails43) AddRepurchaseType() *RepurchaseType11Choice {
	s.RepurchaseType = new(RepurchaseType11Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails43) AddLegalRestrictions() *Restriction3Choice {
	s.LegalRestrictions = new(Restriction3Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails43) AddSecuritiesRTGS() *SecuritiesRTGS3Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS3Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails43) AddSettlingCapacity() *SettlingCapacity3Choice {
	s.SettlingCapacity = new(SettlingCapacity3Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails43) AddSettlementSystemMethod() *SettlementSystemMethod3Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod3Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails43) AddTaxCapacity() *TaxCapacityParty3Choice {
	s.TaxCapacity = new(TaxCapacityParty3Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails43) SetStampDutyIndicator(value string) {
	s.StampDutyIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddStampDutyTaxBasis() *GenericIdentification38 {
	s.StampDutyTaxBasis = new(GenericIdentification38)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails43) AddTracking() *Tracking3Choice {
	s.Tracking = new(Tracking3Choice)
	return s.Tracking
}

func (s *SettlementDetails43) AddAutomaticBorrowing() *AutomaticBorrowing5Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing5Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails43) AddLetterOfGuarantee() *LetterOfGuarantee3Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee3Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails43) SetReturnLeg(value string) {
	s.ReturnLeg = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails43) AddModificationCancellationAllowed() *ModificationCancellationAllowed3Choice {
	s.ModificationCancellationAllowed = new(ModificationCancellationAllowed3Choice)
	return s.ModificationCancellationAllowed
}

func (s *SettlementDetails43) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}
