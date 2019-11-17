package model

// Provides details about the letter of credit or bank guarantee, or other collateral, posted as collateral.
type OtherCollateral6 struct {

	// Identifies the register number of the collateral deposit assigned by the central counterparty.
	AssetNumber *Max35Text `xml:"AsstNb,omitempty"`

	// Provides the unique identification of the letter of credit.
	LetterOfCreditIdentification *Max35Text `xml:"LttrOfCdtId,omitempty"`

	// Amount of the letter/documentary credit.
	LetterOfCreditAmount *ActiveCurrencyAndAmount `xml:"LttrOfCdtAmt,omitempty"`

	// Amount of the bank guarantee.
	GuaranteeAmount *ActiveCurrencyAndAmount `xml:"GrntAmt,omitempty"`

	// Provides a description and an amount of another type of collateral.
	OtherTypeOfCollateral *OtherTypeOfCollateral2 `xml:"OthrTpOfColl,omitempty"`

	// Indicates whether the collateral is proprietarily owned or client owned.
	CollateralOwnership *CollateralOwnership2 `xml:"CollOwnrsh,omitempty"`

	// Date on which the other collateral was issued.
	IssueDate *DateFormat14Choice `xml:"IsseDt,omitempty"`

	// Date on which the other collateral expires.
	ExpiryDate *DateFormat14Choice `xml:"XpryDt,omitempty"`

	// Indicates that the collateral posted in the clearing house covers the margin until a specific timeframe.
	LimitedCoverageIndicator *YesNoIndicator `xml:"LtdCvrgInd,omitempty"`

	// Party that issues the bank guarantee or letter of / documentary credit.
	Issuer *PartyIdentification100Choice `xml:"Issr,omitempty"`

	// Quantity blocked by the central counterparty for any reasonable reason ( for example for judicial reasons). In this case the investor can not withdraw or distribute this collateral.
	BlockedQuantity *FinancialInstrumentQuantity1Choice `xml:"BlckdQty,omitempty"`

	// Valuation date of the other collateral when it was taken as collateral.
	ValueDate *ISODate `xml:"ValDt,omitempty"`

	// Exchange rate.
	ExchangeRate *BaseOneRate `xml:"XchgRate,omitempty"`

	// Value of the collateral based on current market prices.
	MarketValue *ActiveCurrencyAndAmount `xml:"MktVal,omitempty"`

	// Haircut or valuation factor on the collateral expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Value of the collateral after taking into account the haircut, if any.
	CollateralValue *ActiveCurrencyAndAmount `xml:"CollVal"`

	// Place where the securities are safe-kept, physically or notionally.  This place can be, for example, a local custodian, a Central Securities Depository (CSD) or an International Central Securities Depository (ICSD).
	SafekeepingPlace *SafekeepingPlaceFormat10Choice `xml:"SfkpgPlc,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *SecuritiesAccount19 `xml:"SfkpgAcct,omitempty"`
}

func (o *OtherCollateral6) SetAssetNumber(value string) {
	o.AssetNumber = (*Max35Text)(&value)
}

func (o *OtherCollateral6) SetLetterOfCreditIdentification(value string) {
	o.LetterOfCreditIdentification = (*Max35Text)(&value)
}

func (o *OtherCollateral6) SetLetterOfCreditAmount(value, currency string) {
	o.LetterOfCreditAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral6) SetGuaranteeAmount(value, currency string) {
	o.GuaranteeAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral6) AddOtherTypeOfCollateral() *OtherTypeOfCollateral2 {
	o.OtherTypeOfCollateral = new(OtherTypeOfCollateral2)
	return o.OtherTypeOfCollateral
}

func (o *OtherCollateral6) AddCollateralOwnership() *CollateralOwnership2 {
	o.CollateralOwnership = new(CollateralOwnership2)
	return o.CollateralOwnership
}

func (o *OtherCollateral6) AddIssueDate() *DateFormat14Choice {
	o.IssueDate = new(DateFormat14Choice)
	return o.IssueDate
}

func (o *OtherCollateral6) AddExpiryDate() *DateFormat14Choice {
	o.ExpiryDate = new(DateFormat14Choice)
	return o.ExpiryDate
}

func (o *OtherCollateral6) SetLimitedCoverageIndicator(value string) {
	o.LimitedCoverageIndicator = (*YesNoIndicator)(&value)
}

func (o *OtherCollateral6) AddIssuer() *PartyIdentification100Choice {
	o.Issuer = new(PartyIdentification100Choice)
	return o.Issuer
}

func (o *OtherCollateral6) AddBlockedQuantity() *FinancialInstrumentQuantity1Choice {
	o.BlockedQuantity = new(FinancialInstrumentQuantity1Choice)
	return o.BlockedQuantity
}

func (o *OtherCollateral6) SetValueDate(value string) {
	o.ValueDate = (*ISODate)(&value)
}

func (o *OtherCollateral6) SetExchangeRate(value string) {
	o.ExchangeRate = (*BaseOneRate)(&value)
}

func (o *OtherCollateral6) SetMarketValue(value, currency string) {
	o.MarketValue = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral6) SetHaircut(value string) {
	o.Haircut = (*PercentageRate)(&value)
}

func (o *OtherCollateral6) SetCollateralValue(value, currency string) {
	o.CollateralValue = NewActiveCurrencyAndAmount(value, currency)
}

func (o *OtherCollateral6) AddSafekeepingPlace() *SafekeepingPlaceFormat10Choice {
	o.SafekeepingPlace = new(SafekeepingPlaceFormat10Choice)
	return o.SafekeepingPlace
}

func (o *OtherCollateral6) AddSafekeepingAccount() *SecuritiesAccount19 {
	o.SafekeepingAccount = new(SecuritiesAccount19)
	return o.SafekeepingAccount
}
