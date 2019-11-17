package model

// Description of the financial instrument.
type FinancialInstrumentAttributes32 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification14 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat1Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType2Choice `xml:"ClssfctnTp,omitempty"`

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

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *PercentageRate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *PercentageRate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Also known as Minimum Nominal Value. Minimum nominal quantity of financial instrument that must be purchased/sold.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes32) AddFinancialInstrumentIdentification() *SecurityIdentification14 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification14)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes32) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes32) AddDayCountBasis() *InterestComputationMethodFormat1Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat1Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes32) AddClassificationType() *ClassificationType2Choice {
	f.ClassificationType = new(ClassificationType2Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes32) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetPreviousFactor(value string) {
	f.PreviousFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextFactor(value string) {
	f.NextFactor = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes32) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes32) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes32) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes32) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}
