package model

// Parameters which explicitly state the conditions that must be fulfilled before a particular  transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
type SettlementDetails27 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition6Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity1Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification20 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS1Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration1Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership1Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem1Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty1Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType3Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide1Choice `xml:"MktClntSd,omitempty"`

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

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee1Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails27) AddSettlementTransactionCondition() *SettlementTransactionCondition6Choice {
	newValue := new(SettlementTransactionCondition6Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails27) AddSettlingCapacity() *SettlingCapacity1Choice {
	s.SettlingCapacity = new(SettlingCapacity1Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails27) AddStampDutyTaxBasis() *GenericIdentification20 {
	s.StampDutyTaxBasis = new(GenericIdentification20)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails27) AddSecuritiesRTGS() *SecuritiesRTGS1Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS1Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails27) AddRegistration() *Registration1Choice {
	s.Registration = new(Registration1Choice)
	return s.Registration
}

func (s *SettlementDetails27) AddBeneficialOwnership() *BeneficialOwnership1Choice {
	s.BeneficialOwnership = new(BeneficialOwnership1Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails27) AddCashClearingSystem() *CashSettlementSystem1Choice {
	s.CashClearingSystem = new(CashSettlementSystem1Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails27) AddTaxCapacity() *TaxCapacityParty1Choice {
	s.TaxCapacity = new(TaxCapacityParty1Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails27) AddRepurchaseType() *RepurchaseType3Choice {
	s.RepurchaseType = new(RepurchaseType3Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails27) AddMarketClientSide() *MarketClientSide1Choice {
	s.MarketClientSide = new(MarketClientSide1Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails27) AddBlockTrade() *BlockTrade1Choice {
	s.BlockTrade = new(BlockTrade1Choice)
	return s.BlockTrade
}

func (s *SettlementDetails27) AddLegalRestrictions() *Restriction1Choice {
	s.LegalRestrictions = new(Restriction1Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails27) AddSettlementSystemMethod() *SettlementSystemMethod1Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod1Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails27) AddNettingEligibility() *NettingEligibility1Choice {
	s.NettingEligibility = new(NettingEligibility1Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails27) AddCCPEligibility() *CentralCounterPartyEligibility1Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility1Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails27) AddLetterOfGuarantee() *LetterOfGuarantee1Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee1Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails27) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}
