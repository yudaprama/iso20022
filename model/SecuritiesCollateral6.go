package model

// Provides details about the securities posted as collateral.
type SecuritiesCollateral6 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Identification of a security.
	SecurityIdentification *SecurityIdentification19 `xml:"SctyId"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *DateAndDateTimeChoice `xml:"MtrtyDt,omitempty"`

	// Indicates whether the collateral is proprietarily owned or client owned.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Quantity of securities collateral.
	Quantity *FinancialInstrumentQuantity1Choice `xml:"Qty"`

	// Quantity blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedQuantity *FinancialInstrumentQuantity1Choice `xml:"BlckdQty,omitempty"`

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
}

func (s *SecuritiesCollateral6) SetAssetNumber(value string) {
	s.AssetNumber = (*Max35Text)(&value)
}

func (s *SecuritiesCollateral6) AddSecurityIdentification() *SecurityIdentification19 {
	s.SecurityIdentification = new(SecurityIdentification19)
	return s.SecurityIdentification
}

func (s *SecuritiesCollateral6) AddMaturityDate() *DateAndDateTimeChoice {
	s.MaturityDate = new(DateAndDateTimeChoice)
	return s.MaturityDate
}

func (s *SecuritiesCollateral6) AddCollateralOwnership() *CollateralOwnership2 {
	s.CollateralOwnership = new(CollateralOwnership2)
	return s.CollateralOwnership
}

func (s *SecuritiesCollateral6) SetLimitedCoverageIndicator(value string) {
	s.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (s *SecuritiesCollateral6) AddQuantity() *FinancialInstrumentQuantity1Choice {
	s.Quantity = new(FinancialInstrumentQuantity1Choice)
	return s.Quantity
}

func (s *SecuritiesCollateral6) AddBlockedQuantity() *FinancialInstrumentQuantity1Choice {
	s.BlockedQuantity = new(FinancialInstrumentQuantity1Choice)
	return s.BlockedQuantity
}

func (s *SecuritiesCollateral6) AddPrice() *Price2 {
	s.Price = new(Price2)
	return s.Price
}

func (s *SecuritiesCollateral6) SetMarketValue(value, currency string) {
	s.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral6) SetHaircut(value string) {
	s.Haircut = (*PercentageRate)(&value)
}

func (s *SecuritiesCollateral6) SetCollateralValue(value, currency string) {
	s.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SecuritiesCollateral6) SetValueDate(value string) {
	s.ValueDate = (*ISODate)(&value)
}

func (s *SecuritiesCollateral6) AddSafekeepingAccount() *SecuritiesAccount19 {
	s.SafekeepingAccount = new(SecuritiesAccount19)
	return s.SafekeepingAccount
}

func (s *SecuritiesCollateral6) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	s.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return s.SafekeepingPlace
}
