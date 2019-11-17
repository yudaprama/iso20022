package model

// Description of the financial instrument.
type FinancialInstrumentAttributes70 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which an order expires or at which a privilege or offer terminates.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date on which the financial instrument is issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt, for example, in the context of the EU Savings directive.
	PercentageOfDebtClaim *RateFormat3Choice `xml:"PctgOfDebtClm,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Provides the ratio between the quantity of warrants and the quantity of underlying securities.
	WarrantParity *QuantityToQuantityRatio2 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes70) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes70) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes70) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes70) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes70) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes70) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes70) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes70) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes70) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes70) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes70) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes70) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes70) AddWarrantParity() *QuantityToQuantityRatio2 {
	f.WarrantParity = new(QuantityToQuantityRatio2)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes70) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes70) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}
