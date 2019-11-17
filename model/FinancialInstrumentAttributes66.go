package model

// Description of the financial instrument.
type FinancialInstrumentAttributes66 struct {

	// Identifies the financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification19 `xml:"FinInstrmId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

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
	WarrantParity *QuantityToQuantityRatio1 `xml:"WarrtParity,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes66) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes66) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes66) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes66) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes66) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes66) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes66) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes66) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes66) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes66) AddPercentageOfDebtClaim() *RateFormat3Choice {
	f.PercentageOfDebtClaim = new(RateFormat3Choice)
	return f.PercentageOfDebtClaim
}

func (f *FinancialInstrumentAttributes66) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes66) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes66) AddWarrantParity() *QuantityToQuantityRatio1 {
	f.WarrantParity = new(QuantityToQuantityRatio1)
	return f.WarrantParity
}

func (f *FinancialInstrumentAttributes66) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes66) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}
