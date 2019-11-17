package model

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes78 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification4Choice `xml:"PlcOfListg,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat5Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, this is, ownership, of the security.
	RegistrationForm *FormOfSecurity7Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency27Choice `xml:"PmtFrqcy,omitempty"`

	// Status of payment of a security at a particular time.
	PaymentStatus *SecuritiesPaymentStatus6Choice `xml:"PmtSts,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency27Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType33Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle9Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType7Choice `xml:"OptnTp,omitempty"`

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
	CouponAttachedNumber *Number23Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification39 `xml:"PoolNb,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Value of the price, for example, as a currency and value per unit or as a percentage.
	MarketOrIndicativePrice *PriceType2Choice `xml:"MktOrIndctvPric,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price3 `xml:"ExrcPric,omitempty"`

	// Pre-determined price at which the holder of a right is entitled to buy the underlying instrument.
	SubscriptionPrice *Price3 `xml:"SbcptPric,omitempty"`

	// Price of one target security in the conversion.
	ConversionPrice *Price3 `xml:"ConvsPric,omitempty"`

	// Predetermined price at which the holder will have to buy or sell the underlying instrument.
	StrikePrice *Price3 `xml:"StrkPric,omitempty"`

	// Minimum nominal quantity of financial instrument.
	MinimumNominalQuantity *FinancialInstrumentQuantity15Choice `xml:"MinNmnlQty,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity15Choice `xml:"CtrctSz,omitempty"`

	// Identification of the underlying security.
	UnderlyingFinancialInstrumentIdentification []*SecurityIdentification20 `xml:"UndrlygFinInstrmId,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *RestrictedFINXMax350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes78) AddPlaceOfListing() *MarketIdentification4Choice {
	f.PlaceOfListing = new(MarketIdentification4Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes78) AddDayCountBasis() *InterestComputationMethodFormat5Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat5Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes78) AddRegistrationForm() *FormOfSecurity7Choice {
	f.RegistrationForm = new(FormOfSecurity7Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes78) AddPaymentFrequency() *Frequency27Choice {
	f.PaymentFrequency = new(Frequency27Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes78) AddPaymentStatus() *SecuritiesPaymentStatus6Choice {
	f.PaymentStatus = new(SecuritiesPaymentStatus6Choice)
	return f.PaymentStatus
}

func (f *FinancialInstrumentAttributes78) AddVariableRateChangeFrequency() *Frequency27Choice {
	f.VariableRateChangeFrequency = new(Frequency27Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes78) AddClassificationType() *ClassificationType33Choice {
	f.ClassificationType = new(ClassificationType33Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes78) AddOptionStyle() *OptionStyle9Choice {
	f.OptionStyle = new(OptionStyle9Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes78) AddOptionType() *OptionType7Choice {
	f.OptionType = new(OptionType7Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes78) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetYieldToMaturityRate(value string) {
	f.YieldToMaturityRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes78) AddCouponAttachedNumber() *Number23Choice {
	f.CouponAttachedNumber = new(Number23Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes78) AddPoolNumber() *GenericIdentification39 {
	f.PoolNumber = new(GenericIdentification39)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes78) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes78) AddMarketOrIndicativePrice() *PriceType2Choice {
	f.MarketOrIndicativePrice = new(PriceType2Choice)
	return f.MarketOrIndicativePrice
}

func (f *FinancialInstrumentAttributes78) AddExercisePrice() *Price3 {
	f.ExercisePrice = new(Price3)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes78) AddSubscriptionPrice() *Price3 {
	f.SubscriptionPrice = new(Price3)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes78) AddConversionPrice() *Price3 {
	f.ConversionPrice = new(Price3)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes78) AddStrikePrice() *Price3 {
	f.StrikePrice = new(Price3)
	return f.StrikePrice
}

func (f *FinancialInstrumentAttributes78) AddMinimumNominalQuantity() *FinancialInstrumentQuantity15Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity15Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes78) AddContractSize() *FinancialInstrumentQuantity15Choice {
	f.ContractSize = new(FinancialInstrumentQuantity15Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes78) AddUnderlyingFinancialInstrumentIdentification() *SecurityIdentification20 {
	newValue := new(SecurityIdentification20)
	f.UnderlyingFinancialInstrumentIdentification = append(f.UnderlyingFinancialInstrumentIdentification, newValue)
	return newValue
}

func (f *FinancialInstrumentAttributes78) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*RestrictedFINXMax350Text)(&value)
}
