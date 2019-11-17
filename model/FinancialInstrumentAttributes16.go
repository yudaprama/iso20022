package model

// Description of the financial instrument.
type FinancialInstrumentAttributes16 struct {

	// Identification of the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification2 `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised.
	OptionStyle *OptionStyle4Choice `xml:"OptnStyle,omitempty"`

	// Currency in which a financial instrument is currently denominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	NextCouponDate *ISODate `xml:"NxtCpnDt,omitempty"`

	// Date on which the interest rate or redemption price will be/was calculated according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Date on which a financial instrument becomes due and assets are to be repaid.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date on which a financial instrument is called away/redeemed before its scheduled maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date on which a holder of a financial instrument has the right to request redemption of the principal amount prior to its scheduled maturity date.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// Date on which an interest bearing financial instrument begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Deadline by which a convertible security must be converted, according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat19Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes16) AddSecurityIdentification() *SecurityIdentification14 {
	f.SecurityIdentification = new(SecurityIdentification14)
	return f.SecurityIdentification
}

func (f *FinancialInstrumentAttributes16) AddPlaceOfListing() *MarketIdentification2 {
	f.PlaceOfListing = new(MarketIdentification2)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes16) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes16) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes16) AddOptionStyle() *OptionStyle4Choice {
	f.OptionStyle = new(OptionStyle4Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes16) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes16) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes16) AddPreviousFactor() *RateFormat3Choice {
	f.PreviousFactor = new(RateFormat3Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes16) AddNextFactor() *RateFormat3Choice {
	f.NextFactor = new(RateFormat3Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes16) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes16) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes16) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes16) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes16) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes16) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes16) AddIssuePrice() *PriceFormat19Choice {
	f.IssuePrice = new(PriceFormat19Choice)
	return f.IssuePrice
}
