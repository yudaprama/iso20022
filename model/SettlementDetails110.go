package model

// Details of settlement of a transaction.
type SettlementDetails110 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric5Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType24Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade5Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility5Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType17Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility5Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction6Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod5Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing8Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification47 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification47 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails110) AddPriority() *PriorityNumeric5Choice {
	s.Priority = new(PriorityNumeric5Choice)
	return s.Priority
}

func (s *SettlementDetails110) AddSecuritiesTransactionType() *SecuritiesTransactionType24Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType24Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails110) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails110) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails110) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails110) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails110) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails110) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails110) AddExposureType() *ExposureType17Choice {
	s.ExposureType = new(ExposureType17Choice)
	return s.ExposureType
}

func (s *SettlementDetails110) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails110) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails110) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails110) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails110) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails110) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails110) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails110) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails110) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails110) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails110) AddAutomaticBorrowing() *AutomaticBorrowing8Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing8Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails110) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails110) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails110) AddSecuritiesSubBalanceType() *GenericIdentification47 {
	s.SecuritiesSubBalanceType = new(GenericIdentification47)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails110) AddCashSubBalanceType() *GenericIdentification47 {
	s.CashSubBalanceType = new(GenericIdentification47)
	return s.CashSubBalanceType
}
