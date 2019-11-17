package model

// Details of settlement of a transaction.
type SettlementDetails108 struct {

	// Specifies whether the transaction is on hold/blocked/frozen.
	HoldIndicator *YesNoIndicator `xml:"HldInd,omitempty"`

	// Identifies the type of securities transaction.
	SecuritiesTransactionType *SecuritiesTransactionType22Choice `xml:"SctiesTxTp"`

	// Conditions under which the order/trade is to be settled.
	SettlementTransactionCondition []*SettlementTransactionCondition26Choice `xml:"SttlmTxCond,omitempty"`

	// Specifies whether partial settlement is allowed.
	PartialSettlementIndicator *SettlementTransactionCondition5Code `xml:"PrtlSttlmInd,omitempty"`

	// Specifies whether there is change of beneficial ownership.
	BeneficialOwnership *BeneficialOwnership5Choice `xml:"BnfclOwnrsh,omitempty"`

	// Specifies the category of cash clearing system, for example, cheque clearing.
	CashClearingSystem *CashSettlementSystem5Choice `xml:"CshClrSys,omitempty"`

	// Specifies if an instruction is for a market side or a client side transaction.
	MarketClientSide *MarketClientSide5Choice `xml:"MktClntSd,omitempty"`

	// Specifies whether registration should occur upon receipt.
	Registration *Registration11Choice `xml:"Regn,omitempty"`

	// Specifies whether the rate is fixed, variable or a forfeit.
	RepurchaseType *RepurchaseType17Choice `xml:"RpTp,omitempty"`

	// Specifies whether the settlement transaction is to be settled through an RTGS or a non RTGS system.
	SecuritiesRTGS *SecuritiesRTGS5Choice `xml:"SctiesRTGS,omitempty"`

	// Specifies the stamp duty type or exemption reason applicable to the settlement transaction.
	StampDutyTaxBasis *GenericIdentification47 `xml:"StmpDtyTaxBsis,omitempty"`
}

func (s *SettlementDetails108) SetHoldIndicator(value string) {
	s.HoldIndicator = (*YesNoIndicator)(&value)
}

func (s *SettlementDetails108) AddSecuritiesTransactionType() *SecuritiesTransactionType22Choice {
	s.SecuritiesTransactionType = new(SecuritiesTransactionType22Choice)
	return s.SecuritiesTransactionType
}

func (s *SettlementDetails108) AddSettlementTransactionCondition() *SettlementTransactionCondition26Choice {
	newValue := new(SettlementTransactionCondition26Choice)
	s.SettlementTransactionCondition = append(s.SettlementTransactionCondition, newValue)
	return newValue
}

func (s *SettlementDetails108) SetPartialSettlementIndicator(value string) {
	s.PartialSettlementIndicator = (*SettlementTransactionCondition5Code)(&value)
}

func (s *SettlementDetails108) AddBeneficialOwnership() *BeneficialOwnership5Choice {
	s.BeneficialOwnership = new(BeneficialOwnership5Choice)
	return s.BeneficialOwnership
}

func (s *SettlementDetails108) AddCashClearingSystem() *CashSettlementSystem5Choice {
	s.CashClearingSystem = new(CashSettlementSystem5Choice)
	return s.CashClearingSystem
}

func (s *SettlementDetails108) AddMarketClientSide() *MarketClientSide5Choice {
	s.MarketClientSide = new(MarketClientSide5Choice)
	return s.MarketClientSide
}

func (s *SettlementDetails108) AddRegistration() *Registration11Choice {
	s.Registration = new(Registration11Choice)
	return s.Registration
}

func (s *SettlementDetails108) AddRepurchaseType() *RepurchaseType17Choice {
	s.RepurchaseType = new(RepurchaseType17Choice)
	return s.RepurchaseType
}

func (s *SettlementDetails108) AddSecuritiesRTGS() *SecuritiesRTGS5Choice {
	s.SecuritiesRTGS = new(SecuritiesRTGS5Choice)
	return s.SecuritiesRTGS
}

func (s *SettlementDetails108) AddStampDutyTaxBasis() *GenericIdentification47 {
	s.StampDutyTaxBasis = new(GenericIdentification47)
	return s.StampDutyTaxBasis
}
