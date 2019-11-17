package model

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes64 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat4Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity6Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency23Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus5Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency23Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType32Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle8Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType6Choice `xml:"OptnTp,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Next payment date of an interest bearing financial instrument.
	CouponDate *ISODate `xml:"CpnDt,omitempty"`

	// Date on which a privilege expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Date at which the interest rate of an interest bearing security will be calculated and reset, according to the terms of the issue.
	FloatingRateFixingDate *ISODate `xml:"FltgRateFxgDt,omitempty"`

	// Planned final repayment date at the time of issuance.
	MaturityDate *ISODate `xml:"MtrtyDt,omitempty"`

	// Date at which the security was made available.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Next date at which the issuer has the right to pay the security prior to maturity.
	NextCallableDate *ISODate `xml:"NxtCllblDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Rate of return anticipated on a bond when held until maturity date.
	YieldToMaturityRate *PercentageRate `xml:"YldToMtrtyRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number22Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType1Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price2 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price2 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price2 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price2 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity1Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification19 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes64) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes64) AddDayCountBasis() *InterestComputationMethodFormat4Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat4Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes64) AddRegistrationForm() *FormOfSecurity6Choice {
	f.RegistrationForm = new(FormOfSecurity6Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes64) AddPaymentFrequency() *Frequency23Choice {
	f.PaymentFrequency = new(Frequency23Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes64) AddPaymentStatus() *SecuritiesPaymentStatus5Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus5Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes64) AddVariableRateChangeFrequency() *Frequency23Choice {
	f.VariableRateChangeFrequency = new(Frequency23Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes64) AddClassificationType() *ClassificationType32Choice {
	f.ClassificationType = new(ClassificationType32Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes64) AddOptionStyle() *OptionStyle8Choice {
	f.OptionStyle = new(OptionStyle8Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes64) AddOptionType() *OptionType6Choice {
	f.OptionType = new(OptionType6Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes64) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes64) AddCouponAttachedNumber() *Number22Choice {
	f.CouponAttachedNumber = new(Number22Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes64) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes64) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes64) AddMarketOrIndicativePrice() *PriceType1Choice {
	f.MarketOrIndicativePrice = new(PriceType1Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes64) AddExercisePrice() *Price2 {
	f.ExercisePrice = new(Price2)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes64) AddSubscriptionPrice() *Price2 {
	f.SubscriptionPrice = new(Price2)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes64) AddConversionPrice() *Price2 {
	f.ConversionPrice = new(Price2)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes64) AddStrikePrice() *Price2 {
	f.StrikePrice = new(Price2)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes64) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes64) AddContractSize() *FinancialInstrumentQuantity1Choice {
	f.ContractSize = new(FinancialInstrumentQuantity1Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes64) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification19 {
	newValue := new(SecurityIdentification19)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes64) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}
