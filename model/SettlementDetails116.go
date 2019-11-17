package model

// Details of settlement of a transaction.
type SettlementDetails116 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *HoldIndicator7 `xml:"HldInd,omitempty"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition28Choice `xml:"SttlmTxCond,omitempty"`

	// Role of a party in the settlement of the transaction.
	SettlingCapacity *SettlingCapacity8Choice `xml:"SttlgCpcty,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Tax role capacity of the instructing party.
	TaxCapacity *TaxCapacityParty5Choice `xml:"TaxCpcty,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType19Choice `xml:"RpTp,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

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

	// Specifies whether physical settlement may be executed using a letter of guarantee or if the physical certificates should be used.
	LetterOfGuarantee *LetterOfGuarantee5Choice `xml:"LttrOfGrnt,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`
}

func (s *SettlementDetails116) AddHoldIndicator() *HoldIndicator7 {
	s.HoldIndicator = new(HoldIndicator7)
	return s.HoldIndicator
}

func (s *SettlementDetails116) AddSettlementTransactionCondition() *SettlementTransactionCondition28Choice {
	newValue := new(SettlementTransactionCondition28Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails116) AddSettlingCapacity() *SettlingCapacity8Choice {
	s.SettlingCapacity = new(SettlingCapacity8Choice)
	return s.SettlingCapacity
}

func (s *SettlementDetails116) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}

func (s *SettlementDetails116) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails116) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails116) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails116) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails116) AddTaxCapacity() *TaxCapacityParty5Choice {
	s.TaxCapacity = new(TaxCapacityParty5Choice)
	return s.TaxCapacity
}

func (s *SettlementDetails116) AddRepurchaseType() *RepurchaseType19Choice {
	s.RepurchaseType = new(RepurchaseType19Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails116) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails116) AddBlockTrade() *BlockTrade5Choice {
	s.BlockTrade = new(BlockTrade5Choice)
	return s.BlockTrade
}

func (s *SettlementDetails116) AddLegalRestrictions() *Restriction6Choice {
	s.LegalRestrictions = new(Restriction6Choice)
	return s.LegalRestrictions
}

func (s *SettlementDetails116) AddSettlementSystemMethod() *SettlementSystemMethod5Choice {
	s.SettlementSystemMethod = new(SettlementSystemMethod5Choice)
	return s.SettlementSystemMethod
}

func (s *SettlementDetails116) AddNettingEligibility() *NettingEligibility5Choice {
	s.NettingEligibility = new(NettingEligibility5Choice)
	return s.NettingEligibility
}

func (s *SettlementDetails116) AddCCPEligibility() *CentralCounterPartyEligibility5Choice {
	s.CCPEligibility = new(CentralCounterPartyEligibility5Choice)
	return s.CCPEligibility
}

func (s *SettlementDetails116) AddLetterOfGuarantee() *LetterOfGuarantee5Choice {
	s.LetterOfGuarantee = new(LetterOfGuarantee5Choice)
	return s.LetterOfGuarantee
}

func (s *SettlementDetails116) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}
