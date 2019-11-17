package model

// Provides details about the securities posted as collateral.
type SecuritiesCollateral7 struct {

	// Provides the identification of the proposed collateral.
	CollateralIdentification *Max35Text `xml:"CollId,omitempty"`

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Indicates the price of the security.
	Price *Price2 `xml:"Pric,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal,omitempty"`

	// Valuation date of the securities taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc"`

	// Parameters which explicitly state the conditions that must be fulfilled before a particular transaction of a financial instrument can be settled.  These parameters are defined by the instructing party in compliance with settlement rules in the market the transaction will settle in.
	SettlementParameters *SettlementDetails102 `xml:"SttlmParams,omitempty"`
}

func (s *SecuritiesCollateral7) SetCollateralIdentification(value string) {
	s.CollateralIdentification = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral7) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral7) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral7) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral7) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral7) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral7) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral7) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral7) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral7) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral7) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral7) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral7) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}

func (s *SecuritiesCollateral7) AddSettlementParameters() *SettlementDetails102 {
	s.SettlementParameters = new(SettlementDetails102)
	return s.SettlementParameters
}
