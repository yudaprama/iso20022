package model

// Details of settlement of a transaction.
type SettlementDetails90 struct {

	// Specifies whether the transaction was executed with a high priority.
	Priority *PriorityNumeric4Choice `xml:"Prty,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType18Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade was to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement was allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there was change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies whether the settlement instruction was a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Specifies whether the settlement transaction was CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Specifies the underlying business area/type of trade having caused the collateral movement.
	ExposureType *ExposureType16Choice `xml:"XpsrTp,omitempty"`

	// Specifies if an instruction was for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement transaction was eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether registration was to occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies the type of repurchase transaction.
	RepurchaseType *RepurchaseType12Choice `xml:"RpTp,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement transaction was to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies whether the settlement instruction was to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Condition for automatic borrowing.
	AutomaticBorrowing *AutomaticBorrowing6Choice `xml:"AutomtcBrrwg,omitempty"`

	// Specifies whether physical settlement was executed using a letter of guarantee or if the physical certificates were used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether securities were requested to be included in the pool of securities eligible for collateral purposes.
	EligibleForCollateral *YesNoIndicator `xml:"ElgblForColl,omitempty"`

	// Specifies the securities sub balance type indicator (example restriction type for a market infrastructure).
	SecuritiesSubBalanceType *GenericIdentification30 `xml:"SctiesSubBalTp,omitempty"`

	// Specifies the cash sub balance type indicator (example restriction type for a market infrastructure).
	CashSubBalanceType *GenericIdentification30 `xml:"CshSubBalTp,omitempty"`
}

func (s *SettlementDetails90) AddPriority() *PriorityNumeric4Choice {
	s.Priority = new(PriorityNumeric4Choice)
	return s.Priority
}

func (s *SettlementDetails90) AddSecuritiesTransactionType() *SecuritiesTransactionType18Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType18Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails90) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails90) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails90) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails90) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails90) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails90) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails90) AddExposureType() *ExposureType16Choice {
	s.ExposureType = new(ExposureType16Choice)
	return s.ExposureType
}

func (s *SettlementDetails90) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails90) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails90) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails90) AddRepurchaseType() *RepurchaseType12Choice {
	s.RepurchaseType = new(RepurchaseType12Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails90) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails90) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails90) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails90) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails90) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails90) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails90) AddAutomaticBorrowing() *AutomaticBorrowing6Choice {
	s.AutomaticBorrowing = new(AutomaticBorrowing6Choice)
	return s.AutomaticBorrowing
}

func (s *SettlementDetails90) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails90) SetEligibleForCollateral(value string) {
	s.EligibleForCollateral = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails90) AddSecuritiesSubBalanceType() *GenericIdentification30 {
	s.SecuritiesSubBalanceType = new(GenericIdentification30)
	return s.SecuritiesSubBalanceType
}

func (s *SettlementDetails90) AddCashSubBalanceType() *GenericIdentification30 {
	s.CashSubBalanceType = new(GenericIdentification30)
	return s.CashSubBalanceType
}
