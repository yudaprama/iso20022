package model

// Additional restrictions on the financial instrument, related to the stipulation.
type FinancialInstrumentStipulations2 struct {

	// Type of stipulation expressing geographical constraints on a fixed income instrument. It is expressed with a state or country abbreviation and a minimum or maximum percentage. Example: CA 0-80 (minimum of 80 percent in Californian assests).
	Geographics *Max35Text `xml:"Geogcs,omitempty"`

	// Range of allowed yield.
	YieldRange *AmountOrPercentageRange `xml:"YldRg,omitempty"`

	// Range of assessment of securities credit and investment risk.
	Rating *Rating1 `xml:"Ratg,omitempty"`

	// Identification of a range of coupon numbers attached to its related financial instrument.
	CouponRange *AmountOrPercentageRange `xml:"CpnRg,omitempty"`

	// Indicates whether the financial instrument repays the principal amount in parts during the life cycle of the security.
	AmortisableIndicator *YesNoIndicator `xml:"AmtsblInd,omitempty"`

	// Reason for which money is raised through the issuance of a security.
	Purpose *Max256Text `xml:"Purp,omitempty"`

	// Identifies whether the issue is subject to alternative minimum taxation (used for municipal bonds).
	AlternativeMinimumTaxIndicator *YesNoIndicator `xml:"AltrntvMinTaxInd,omitempty"`

	// Indicates an instruction to reinvest dividends in the underlying security (or proceeds at maturity in a similar instrument) if the current rate is <rate> or better.
	AutoReinvestment *PercentageRate `xml:"AutoRinvstmt,omitempty"`

	// Indicates the conditions under which the order/trade is to be/was executed.
	TransactionConditions *TradeTransactionCondition2Code `xml:"TxConds,omitempty"`

	// Currency in which a security is issued or redenominated.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Indicates an instruction to override an investment's default start and/or end date with a custom date.
	CustomDate *DateTimePeriodDetails1 `xml:"CstmDt,omitempty"`

	// Haircut or valuation factor on the security expressed as a percentage.
	Haircut *PercentageRate `xml:"Hrcut,omitempty"`

	// Identifies whether the lender is assured partial or full payment by a third party if the borrower defaults.
	InsuredIndicator *YesNoIndicator `xml:"InsrdInd,omitempty"`

	// Indicates an instruction or attribute giving the number of days to be included in the look-back period for the investment. E.g. some options allow exercise based on the underlying asset's optimal value over the look-back period.
	LookBack *Number `xml:"LookBck,omitempty"`

	// Indicates the maturity date.
	MaturityDate *ISOYearMonth `xml:"MtrtyDt,omitempty"`

	// Indicates the issue date.
	IssueDate *ISOYearMonth `xml:"IsseDt,omitempty"`

	// Identification of the issuer.
	IssuerIdentification *BICNonFIIdentifier `xml:"IssrId,omitempty"`

	// Identifies the issue size range.
	IssueSize *Number `xml:"IsseSz,omitempty"`

	// Indicates the minimum denomination of a security.
	MinimumDenomination *FinancialInstrumentQuantityChoice `xml:"MinDnmtn,omitempty"`

	// Maximum number of time the collateral can be substitute.
	MaximumSubstitution *Number `xml:"MaxSbstitn,omitempty"`

	// Indicates the minimum tradable increments of a security.
	MinimumIncrement *FinancialInstrumentQuantityChoice `xml:"MinIncrmt,omitempty"`

	// Indicates the periodic or regular cycle of interest payments.
	PaymentFrequency *Frequency1Code `xml:"PmtFrqcy,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumQuantity *FinancialInstrumentQuantityChoice `xml:"MinQty,omitempty"`

	// Indicates a search criterion used when looking to buy a bond, particularly an MBS, issued in a particular year.
	Production *Max35Text `xml:"Pdctn,omitempty"`

	// Identifies if the securities is restricted or not (as per Rule 144).
	RestrictedIndicator *YesNoIndicator `xml:"RstrctdInd,omitempty"`

	// Indicates the frequency at which the bond is re-rated and therefore re-priced (bond attribute, particularly of floating rate and index linked instruments).
	PriceFrequency *Frequency1Code `xml:"PricFrqcy,omitempty"`

	// Indicates the market sector the security is classified as.  E.g. pharmacuticals, automobile, housing, etc.
	Sector *Max35Text `xml:"Sctr,omitempty"`

	// Indicates the maximum number of times collateral can be substituted.
	SubstitutionFrequency *Frequency1Code `xml:"SbstitnFrqcy,omitempty"`

	// Number of remaining times the collateral can be substitute.
	SubstitutionLeft *Number `xml:"SbstitnLft,omitempty"`

	// Indicates a search criterion when looking to buy an MBS that either is [yes] or is not [no] an entire pool.
	WholePoolIndicator *YesNoIndicator `xml:"WhlPoolInd,omitempty"`

	// Identifies the Benchmark source price (eg. BB Generic, BB Fairvalue, Brokertec..).
	PriceSource *Max35Text `xml:"PricSrc,omitempty"`

	// Date/time at which an interest bearing security becomes due and assets are to be repaid.
	ExpirationDate *ISODateTime `xml:"XprtnDt,omitempty"`

	// Amount for which a security can be overalloted (as in greenshoe option).
	OverAllotmentAmount *ActiveCurrencyAndAmount `xml:"OverAlltmtAmt,omitempty"`

	// Percentage for which a security can be overalloted (as in greenshoe option).
	OverAllotmentRate *PercentageRate `xml:"OverAlltmtRate,omitempty"`

	// Indicates a search criterion used when looking to buy a bond within a particular price range.
	PriceRange *AmountOrPercentageRange `xml:"PricRg,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the interest bearing security is convertible into another type of security.
	ConvertibleIndicator *YesNoIndicator `xml:"ConvtblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Indicates whether an interest bearing instrument is deposited in a fund that will be used to pay debt service on refunded securities.
	PreFundedIndicator *YesNoIndicator `xml:"PreFnddInd,omitempty"`

	// Indicates whether an interest bearing instrument is being escrowed or collateralized either by direct obligations guaranteed by the US government, or by other types of securities. The maturity schedules of the securities in the escrow fund are determined in such a way to pay the maturity value, coupon, and premium payments (if any) of the refunded bonds.
	EscrowedIndicator *YesNoIndicator `xml:"EscrwdInd,omitempty"`

	// Indicates whether the security has no maturity date.
	PerpetualIndicator *YesNoIndicator `xml:"PerptlInd,omitempty"`
}

func (f *FinancialInstrumentStipulations2) SetGeographics(value string) {
	f.Geographics = (*Max35Text)(&value)
}

func (f *FinancialInstrumentStipulations2) AddYieldRange() *AmountOrPercentageRange {
	f.YieldRange = new(AmountOrPercentageRange)
	return f.YieldRange
}

func (f *FinancialInstrumentStipulations2) AddRating() *Rating1 {
	f.Rating = new(Rating1)
	return f.Rating
}

func (f *FinancialInstrumentStipulations2) AddCouponRange() *AmountOrPercentageRange {
	f.CouponRange = new(AmountOrPercentageRange)
	return f.CouponRange
}

func (f *FinancialInstrumentStipulations2) SetAmortisableIndicator(value string) {
	f.AmortisableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPurpose(value string) {
	f.Purpose = (*Max256Text)(&value)
}

func (f *FinancialInstrumentStipulations2) SetAlternativeMinimumTaxIndicator(value string) {
	f.AlternativeMinimumTaxIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetAutoReinvestment(value string) {
	f.AutoReinvestment = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentStipulations2) SetTransactionConditions(value string) {
	f.TransactionConditions = (*TradeTransactionCondition2Code)(&value)
}

func (f *FinancialInstrumentStipulations2) SetCurrency(value string) {
	f.Currency = (*CurrencyCode)(&value)
}

func (f *FinancialInstrumentStipulations2) AddCustomDate() *DateTimePeriodDetails1 {
	f.CustomDate = new(DateTimePeriodDetails1)
	return f.CustomDate
}

func (f *FinancialInstrumentStipulations2) SetHaircut(value string) {
	f.Haircut = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentStipulations2) SetInsuredIndicator(value string) {
	f.InsuredIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetLookBack(value string) {
	f.LookBack = (*Number)(&value)
}

func (f *FinancialInstrumentStipulations2) SetMaturityDate(value string) {
	f.MaturityDate = (*ISOYearMonth)(&value)
}

func (f *FinancialInstrumentStipulations2) SetIssueDate(value string) {
	f.IssueDate = (*ISOYearMonth)(&value)
}

func (f *FinancialInstrumentStipulations2) SetIssuerIdentification(value string) {
	f.IssuerIdentification = (*BICNonFIIdentifier)(&value)
}

func (f *FinancialInstrumentStipulations2) SetIssueSize(value string) {
	f.IssueSize = (*Number)(&value)
}

func (f *FinancialInstrumentStipulations2) AddMinimumDenomination() *FinancialInstrumentQuantityChoice {
	f.MinimumDenomination = new(FinancialInstrumentQuantityChoice)
	return f.MinimumDenomination
}

func (f *FinancialInstrumentStipulations2) SetMaximumSubstitution(value string) {
	f.MaximumSubstitution = (*Number)(&value)
}

func (f *FinancialInstrumentStipulations2) AddMinimumIncrement() *FinancialInstrumentQuantityChoice {
	f.MinimumIncrement = new(FinancialInstrumentQuantityChoice)
	return f.MinimumIncrement
}

func (f *FinancialInstrumentStipulations2) SetPaymentFrequency(value string) {
	f.PaymentFrequency = (*Frequency1Code)(&value)
}

func (f *FinancialInstrumentStipulations2) AddMinimumQuantity() *FinancialInstrumentQuantityChoice {
	f.MinimumQuantity = new(FinancialInstrumentQuantityChoice)
	return f.MinimumQuantity
}

func (f *FinancialInstrumentStipulations2) SetProduction(value string) {
	f.Production = (*Max35Text)(&value)
}

func (f *FinancialInstrumentStipulations2) SetRestrictedIndicator(value string) {
	f.RestrictedIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPriceFrequency(value string) {
	f.PriceFrequency = (*Frequency1Code)(&value)
}

func (f *FinancialInstrumentStipulations2) SetSector(value string) {
	f.Sector = (*Max35Text)(&value)
}

func (f *FinancialInstrumentStipulations2) SetSubstitutionFrequency(value string) {
	f.SubstitutionFrequency = (*Frequency1Code)(&value)
}

func (f *FinancialInstrumentStipulations2) SetSubstitutionLeft(value string) {
	f.SubstitutionLeft = (*Number)(&value)
}

func (f *FinancialInstrumentStipulations2) SetWholePoolIndicator(value string) {
	f.WholePoolIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPriceSource(value string) {
	f.PriceSource = (*Max35Text)(&value)
}

func (f *FinancialInstrumentStipulations2) SetExpirationDate(value string) {
	f.ExpirationDate = (*ISODateTime)(&value)
}

func (f *FinancialInstrumentStipulations2) SetOverAllotmentAmount(value, currency string) {
	f.OverAllotmentAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentStipulations2) SetOverAllotmentRate(value string) {
	f.OverAllotmentRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentStipulations2) AddPriceRange() *AmountOrPercentageRange {
	f.PriceRange = new(AmountOrPercentageRange)
	return f.PriceRange
}

func (f *FinancialInstrumentStipulations2) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetConvertibleIndicator(value string) {
	f.ConvertibleIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPreFundedIndicator(value string) {
	f.PreFundedIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetEscrowedIndicator(value string) {
	f.EscrowedIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentStipulations2) SetPerpetualIndicator(value string) {
	f.PerpetualIndicator = (*YesNoIndicator)(&value)
}
