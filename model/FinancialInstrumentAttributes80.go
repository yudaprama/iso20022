package model

// Description of the financial instrument.
type FinancialInstrumentAttributes80 struct {

	// Identification of the financial instrument.
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
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Annualised interest rate of a financial instrument used to calculate the actual interest rate of the coupon or the accrued interest.
	InterestRate *RateFormat3Choice `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *RateFormat3Choice `xml:"NxtIntrstRate,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument that may be instructed.
	MinimumQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinQtyToInst,omitempty"`

	// Minimum multiple quantity of financial instrument that may be instructed.
	MinimumMultipleQuantityToInstruct *FinancialInstrumentQuantity1Choice `xml:"MinMltplQtyToInst,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a financial instrument quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Initial issue/valuation price of a resulting security under a corporate action.
	IssuePrice *PriceFormat45Choice `xml:"IssePric,omitempty"`
}

func (f *FinancialInstrumentAttributes80) AddFinancialInstrumentIdentification() *SecurityIdentification19 {
	f.FinancialInstrumentIdentification = new(SecurityIdentification19)
	return f.FinancialInstrumentIdentification
}

func (f *FinancialInstrumentAttributes80) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes80) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes80) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes80) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes80) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes80) SetNextCouponDate(value string) {
	f.NextCouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes80) AddPreviousFactor() *RateFormat12Choice {
	f.PreviousFactor = new(RateFormat12Choice)
	return f.PreviousFactor
}

func (f *FinancialInstrumentAttributes80) AddNextFactor() *RateFormat12Choice {
	f.NextFactor = new(RateFormat12Choice)
	return f.NextFactor
}

func (f *FinancialInstrumentAttributes80) AddInterestRate() *RateFormat3Choice {
	f.InterestRate = new(RateFormat3Choice)
	return f.InterestRate
}

func (f *FinancialInstrumentAttributes80) AddNextInterestRate() *RateFormat3Choice {
	f.NextInterestRate = new(RateFormat3Choice)
	return f.NextInterestRate
}

func (f *FinancialInstrumentAttributes80) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes80) AddMinimumQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumQuantityToInstruct
}

func (f *FinancialInstrumentAttributes80) AddMinimumMultipleQuantityToInstruct() *FinancialInstrumentQuantity1Choice {
	f.MinimumMultipleQuantityToInstruct = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumMultipleQuantityToInstruct
}

func (f *FinancialInstrumentAttributes80) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes80) AddIssuePrice() *PriceFormat45Choice {
	f.IssuePrice = new(PriceFormat45Choice)
	return f.IssuePrice
}
