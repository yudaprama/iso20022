package model

// Description of the financial instrument.
type FinancialInstrumentAttributes69 struct {

	// Identification of a financial instrument.
	FinancialInstrumentIdentification *SecurityIdentification20 `xml:"FinInstrmId,omitempty"`

	// Place where the referenced financial instrument is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the financial instrument.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI).
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

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
	PreviousFactor *BaseOne14Rate `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *BaseOne14Rate `xml:"NxtFctr,omitempty"`

	// Annual rate of a financial instrument.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`
}

func (f *FinancialInstrumentAttributes69) AddFinancialInstrumentIdentification() *SecurityIdentification20 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification20)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes69) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes69) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes69) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes69) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextFactor(value string) {
	f.NextFactor = (*BaseOne14Rate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes69) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes69) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes69) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}
