package model

// Details of settlement of a transaction.
type SettlementDetails126 struct {

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition16Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity7Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification30 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS4Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration9Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership4Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem4Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty4Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType22Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide4Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether the settlement instruction is a block parent or child.
	BlockTrade *BlockTrade4Choice `xml:"BlckTrad,omitempty"`

	// Regulatory restrictions applicable to a security.
	LegalRestrictions *Restriction5Choice `xml:"LglRstrctns,omitempty"`

	// Specifies whether the settlement instruction is to be settled through the default or the alternate settlement system.
	SettlementSystemMethod *SettlementSystemMethod4Choice `xml:"SttlmSysMtd,omitempty"`

	// Specifies whether the settlement transaction is eligible for netting.
	NettingEligibility *NettingEligibility4Choice `xml:"NetgElgblty,omitempty"`

	// Specifies whether the settlement transaction is CCP (Central Counterparty) eligible.
	CCPEligibility *CentralCounterPartyEligibility4Choice `xml:"CCPElgblty,omitempty"`

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee4Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails126) AddSettlementTransactionCondition() *SettlementTransactionCondition16Choice {
	newValue := new(SettlementTransactionCondition16Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails126) AddSettlingCapacity() *SettlingCapacity7Choice {
	s.SettlingCapacity = new(SettlingCapacity7Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails126) AddStampDutyTaxBasis() *GenericIdentification30 {
	s.StampDutyTaxBasis = new(GenericIdentification30)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails126) AddSecuritiesRTGS() *SecuritiesRTGS4Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS4Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails126) AddRegistration() *Registration9Choice {
	s.Registration = new(Registration9Choice)
	return s.Registration
}

func (s *SettlementDetails126) AddBeneficialOwnership() *BeneficialOwnership4Choice {
	s.BeneficialOwnership = new(BeneficialOwnership4Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails126) AddCashClearingSystem() *CashSettlementSystem4Choice {
	s.CashClearingSystem = new(CashSettlementSystem4Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails126) AddTaxCapacity() *TaxCapacityParty4Choice {
	s.TaxCapacity = new(TaxCapacityParty4Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails126) AddRepurchaseType() *RepurchaseType22Choice {
	s.RepurchaseType = new(RepurchaseType22Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails126) AddMarketClientSide() *MarketClientSide4Choice {
	s.MarketClientSide = new(MarketClientSide4Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails126) AddBlockTrade() *BlockTrade4Choice {
	s.BlockTrade = new(BlockTrade4Choice)
	return s.BlockTrade
}

func (s *SettlementDetails126) AddLegalRestrictions() *Restriction5Choice {
	s.LegalRestrictions = new(Restriction5Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails126) AddSettlementSystemMethod() *SettlementSystemMethod4Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod4Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails126) AddNettingEligibility() *NettingEligibility4Choice {
	s.NettingEligibility = new(NettingEligibility4Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails126) AddCCPEligibility() *CentralCounterPartyEligibility4Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility4Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails126) AddLetterOfGuarantee() *LetterOfGuarantee4Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee4Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails126) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}
