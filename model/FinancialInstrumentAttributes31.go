package model

// Elements characterising a financial instrument.
type FinancialInstrumentAttributes31 struct {

	// Market(s) on which the security is listed.
	PlaceOfListing *MarketIdentification3Choice `xml:"PlcOfListg,omitempty"`

	// Assessment of securities credit and investment risk.
	Rating *Rating1 `xml:"Ratg,omitempty"`

	// Unique and unambiguous identifier of a certificate assigned by the issuer.
	CertificateNumber *Max35Text `xml:"CertNb,omitempty"`

	// Specifies the computation method of (accrued) interest of the security.
	DayCountBasis *InterestComputationMethodFormat3Choice `xml:"DayCntBsis,omitempty"`

	// Specifies the form, ie, ownership, of the security.
	RegistrationForm *FormOfSecurity4Choice `xml:"RegnForm,omitempty"`

	// Specifies the frequency of an interest payment.
	PaymentFrequency *Frequency7Choice `xml:"PmtFrqcy,omitempty"`

	// Specifies the frequency of change to the variable rate of an interest bearing instrument.
	VariableRateChangeFrequency *Frequency7Choice `xml:"VarblRateChngFrqcy,omitempty"`

	// Classification type of the financial instrument, as per the ISO Classification of Financial Instrument (CFI) codification, for example, common share with voting rights, fully paid, or registered.
	ClassificationType *ClassificationType30Choice `xml:"ClssfctnTp,omitempty"`

	// Specifies how an option can be exercised (American, European, Bermudan)
	OptionStyle *OptionStyle6Choice `xml:"OptnStyle,omitempty"`

	// Specifies whether it is a Call option (right to purchase a specific underlying asset) or a Put option (right to sell a specific underlying asset).
	OptionType *OptionType4Choice `xml:"OptnTp,omitempty"`

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

	// Deadline by which a convertible security must be converted according to the terms of the issue.
	ConversionDate *ISODate `xml:"ConvsDt,omitempty"`

	// Date at which the holder has the right to ask for redemption of the security prior to final maturity.
	PutableDate *ISODate `xml:"PutblDt,omitempty"`

	// First date at which a security begins to accrue interest.
	DatedDate *ISODate `xml:"DtdDt,omitempty"`

	// Date at which the first interest payment is due to holders of the security.
	FirstPaymentDate *ISODate `xml:"FrstPmtDt,omitempty"`

	// Date on which the current factor will be changed to the next factor.
	NextFactorDate *ISODate `xml:"NxtFctrDt,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that was applicable before the current factor and defines the outstanding principal of the financial instrument (for factored securities).
	PreviousFactor *BaseOneRate `xml:"PrvsFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 defining the outstanding principal of the financial instrument (for factored securities).
	CurrentFactor *BaseOneRate `xml:"CurFctr,omitempty"`

	// Rate expressed as a decimal between 0 and 1 that will be applicable as of the next factor date and defines the outstanding principal of the financial instrument (for factored securities).
	NextFactor *BaseOneRate `xml:"NxtFctr,omitempty"`

	// End ratio of principal outstanding to the original balance.
	EndFactor *BaseOneRate `xml:"EndFctr,omitempty"`

	// Per annum ratio of interest paid to the principal amount of the financial instrument for a specific period of time.
	InterestRate *PercentageRate `xml:"IntrstRate,omitempty"`

	// Interest rate applicable to the next interest payment period in relation to variable rate instruments.
	NextInterestRate *PercentageRate `xml:"NxtIntrstRate,omitempty"`

	// Specifies the reference rate for fixed income instruments where the
	// price of the instrument is indexed to the price of an underlying benchmark.
	IndexRateBasis *PercentageRate `xml:"IndxRateBsis,omitempty"`

	// Percentage of the underlying assets of a fund that represents a debt and is in the scope of the EU Savings directive.
	PercentageOfDebtClaims *PercentageRate `xml:"PctgOfDebtClms,omitempty"`

	// Number of the coupon attached to the physical security.
	CouponAttachedNumber *Number1Choice `xml:"CpnAttchdNb,omitempty"`

	// Number identifying a group of underlying assets assigned by the issuer of a factored security.
	PoolNumber *GenericIdentification37 `xml:"PoolNb,omitempty"`

	// Number allocated by options exchanges to record that an option has undergone a change in its contract specifications (particularly adjustment of the strike price)
	VersionNumber *Number1Choice `xml:"VrsnNb,omitempty"`

	// Indicates whether the interest bearing security is convertible into another type of security.
	ConvertibleIndicator *YesNoIndicator `xml:"ConvtblInd,omitempty"`

	// Indicates whether the interest rate of an interest bearing instrument is reset periodically.
	VariableRateIndicator *YesNoIndicator `xml:"VarblRateInd,omitempty"`

	// Indicates whether the underlying security is owned by the writer of the entitlement.
	CoveredIndicator *YesNoIndicator `xml:"CvrdInd,omitempty"`

	// Indicates whether the issuer has the right to pay the security prior to maturity. Also called RetractableIndicator.
	CallableIndicator *YesNoIndicator `xml:"CllblInd,omitempty"`

	// Indicates whether the holder has the right to ask for redemption of the security prior to final maturity. Also called RedeemableIndicator.
	PutableIndicator *YesNoIndicator `xml:"PutblInd,omitempty"`

	// Indicates whether the warrants on a financial instrument (which has been traded cum warrants) will be attached on delivery.
	WarrantAttachedOnDelivery *YesNoIndicator `xml:"WarrtAttchdOnDlvry,omitempty"`

	// Indicates whether the payment of the coupon (interest) on an interest bearing instrument is off the normal schedule.
	OddCouponIndicator *YesNoIndicator `xml:"OddCpnInd,omitempty"`

	// Indicates whether, in the case of a debt security subject to redemption before maturity, such redemption could have an impact on the yield.
	RedemptionYieldImpact *YesNoIndicator `xml:"RedYldImpct,omitempty"`

	// Indicates whether the actual yield of an asset-backed security may vary according to the rate at which the underlying receivables or other financial assets are prepaid.
	YieldVariance *YesNoIndicator `xml:"YldVar,omitempty"`

	// Predetermined price at which the holder of a derivative will buy or sell the underlying instrument.
	ExercisePrice *Price4 `xml:"ExrcPric,omitempty"`

	// Fixed price at which a new securities issue is being offered to the public.
	SubscriptionPrice *Price4 `xml:"SbcptPric,omitempty"`

	// Price at which a given convertible security can be converted to common stock .
	ConversionPrice *Price4 `xml:"ConvsPric,omitempty"`

	// Amount included in the Net Asset Value(NAV) that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerShare *Price4 `xml:"TaxblIncmPerShr,omitempty"`

	// Indicates the minimum tradable quantity of a security.
	MinimumNominalQuantity *FinancialInstrumentQuantity1Choice `xml:"MinNmnlQty,omitempty"`

	// Minimum quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblQty,omitempty"`

	// Minimum multiple quantity of financial instrument or lot of rights/warrants that must be exercised.
	MinimumExercisableMultipleQuantity *FinancialInstrumentQuantity1Choice `xml:"MinExrcblMltplQty,omitempty"`

	// Signed face amount and amortised value of security.
	FaceAmount *ImpliedCurrencyAndAmount `xml:"FaceAmt,omitempty"`

	// Ratio or multiplying factor used to convert one contract into a quantity.
	ContractSize *FinancialInstrumentQuantity18Choice `xml:"CtrctSz,omitempty"`

	// Provides additional information about the financial instrument in narrative form.
	FinancialInstrumentAttributeAdditionalDetails *Max350Text `xml:"FinInstrmAttrAddtlDtls,omitempty"`
}

func (f *FinancialInstrumentAttributes31) AddPlaceOfListing() *MarketIdentification3Choice {
	f.PlaceOfListing = new(MarketIdentification3Choice)
	return f.PlaceOfListing
}

func (f *FinancialInstrumentAttributes31) AddRating() *Rating1 {
	f.Rating = new(Rating1)
	return f.Rating
}

func (f *FinancialInstrumentAttributes31) SetCertificateNumber(value string) {
	f.CertificateNumber = (*Max35Text)(&value)
}

func (f *FinancialInstrumentAttributes31) AddDayCountBasis() *InterestComputationMethodFormat3Choice {
	f.DayCountBasis = new(InterestComputationMethodFormat3Choice)
	return f.DayCountBasis
}

func (f *FinancialInstrumentAttributes31) AddRegistrationForm() *FormOfSecurity4Choice {
	f.RegistrationForm = new(FormOfSecurity4Choice)
	return f.RegistrationForm
}

func (f *FinancialInstrumentAttributes31) AddPaymentFrequency() *Frequency7Choice {
	f.PaymentFrequency = new(Frequency7Choice)
	return f.PaymentFrequency
}

func (f *FinancialInstrumentAttributes31) AddVariableRateChangeFrequency() *Frequency7Choice {
	f.VariableRateChangeFrequency = new(Frequency7Choice)
	return f.VariableRateChangeFrequency
}

func (f *FinancialInstrumentAttributes31) AddClassificationType() *ClassificationType30Choice {
	f.ClassificationType = new(ClassificationType30Choice)
	return f.ClassificationType
}

func (f *FinancialInstrumentAttributes31) AddOptionStyle() *OptionStyle6Choice {
	f.OptionStyle = new(OptionStyle6Choice)
	return f.OptionStyle
}

func (f *FinancialInstrumentAttributes31) AddOptionType() *OptionType4Choice {
	f.OptionType = new(OptionType4Choice)
	return f.OptionType
}

func (f *FinancialInstrumentAttributes31) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCouponDate(value string) {
	f.CouponDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetExpiryDate(value string) {
	f.ExpiryDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetFloatingRateFixingDate(value string) {
	f.FloatingRateFixingDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetMaturityDate(value string) {
	f.MaturityDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetIssueDate(value string) {
	f.IssueDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextCallableDate(value string) {
	f.NextCallableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetConversionDate(value string) {
	f.ConversionDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPutableDate(value string) {
	f.PutableDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetDatedDate(value string) {
	f.DatedDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetFirstPaymentDate(value string) {
	f.FirstPaymentDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextFactorDate(value string) {
	f.NextFactorDate = (*ISODate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPreviousFactor(value string) {
	f.PreviousFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCurrentFactor(value string) {
	f.CurrentFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextFactor(value string) {
	f.NextFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetEndFactor(value string) {
	f.EndFactor = (*BaseOneRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetInterestRate(value string) {
	f.InterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetNextInterestRate(value string) {
	f.NextInterestRate = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetIndexRateBasis(value string) {
	f.IndexRateBasis = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPercentageOfDebtClaims(value string) {
	f.PercentageOfDebtClaims = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentAttributes31) AddCouponAttachedNumber() *Number1Choice {
	f.CouponAttachedNumber = new(Number1Choice)
	return f.CouponAttachedNumber
}

func (f *FinancialInstrumentAttributes31) AddPoolNumber() *GenericIdentification37 {
	f.PoolNumber = new(GenericIdentification37)
	return f.PoolNumber
}

func (f *FinancialInstrumentAttributes31) AddVersionNumber() *Number1Choice {
	f.VersionNumber = new(Number1Choice)
	return f.VersionNumber
}

func (f *FinancialInstrumentAttributes31) SetConvertibleIndicator(value string) {
	f.ConvertibleIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetVariableRateIndicator(value string) {
	f.VariableRateIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCoveredIndicator(value string) {
	f.CoveredIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetCallableIndicator(value string) {
	f.CallableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetPutableIndicator(value string) {
	f.PutableIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetWarrantAttachedOnDelivery(value string) {
	f.WarrantAttachedOnDelivery = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetOddCouponIndicator(value string) {
	f.OddCouponIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetRedemptionYieldImpact(value string) {
	f.RedemptionYieldImpact = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) SetYieldVariance(value string) {
	f.YieldVariance = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrumentAttributes31) AddExercisePrice() *Price4 {
	f.ExercisePrice = new(Price4)
	return f.ExercisePrice
}

func (f *FinancialInstrumentAttributes31) AddSubscriptionPrice() *Price4 {
	f.SubscriptionPrice = new(Price4)
	return f.SubscriptionPrice
}

func (f *FinancialInstrumentAttributes31) AddConversionPrice() *Price4 {
	f.ConversionPrice = new(Price4)
	return f.ConversionPrice
}

func (f *FinancialInstrumentAttributes31) AddTaxableIncomePerShare() *Price4 {
	f.TaxableIncomePerShare = new(Price4)
	return f.TaxableIncomePerShare
}

func (f *FinancialInstrumentAttributes31) AddMinimumNominalQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumNominalQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumNominalQuantity
}

func (f *FinancialInstrumentAttributes31) AddMinimumExercisableQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableQuantity
}

func (f *FinancialInstrumentAttributes31) AddMinimumExercisableMultipleQuantity() *FinancialInstrumentQuantity1Choice {
	f.MinimumExercisableMultipleQuantity = new(FinancialInstrumentQuantity1Choice)
	return f.MinimumExercisableMultipleQuantity
}

func (f *FinancialInstrumentAttributes31) SetFaceAmount(value, currency string) {
	f.FaceAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentAttributes31) AddContractSize() *FinancialInstrumentQuantity18Choice {
	f.ContractSize = new(FinancialInstrumentQuantity18Choice)
	return f.ContractSize
}

func (f *FinancialInstrumentAttributes31) SetFinancialInstrumentAttributeAdditionalDetails(value string) {
	f.FinancialInstrumentAttributeAdditionalDetails = (*Max350Text)(&value)
}
