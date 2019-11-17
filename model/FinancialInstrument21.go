package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument21 struct {

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Specifies the form, that is, ownership, of the security.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Company specific description of a group of funds.
	ProductGroup *Max140Text `xml:"PdctGrp,omitempty"`

	// Name of the umbrella fund in which financial instrument is contained.
	UmbrellaName *Max35Text `xml:"UmbrllNm,omitempty"`

	// Currency of the investment fund class.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd,omitempty"`

	// Country where the fund has legal domicile as reflected in the ISIN classification.
	CountryOfDomicile *CountryCode `xml:"CtryOfDmcl,omitempty"`

	// Countries where the fund is registered for distribution.
	RegisteredDistributionCountry []*CountryCode `xml:"RegdDstrbtnCtry,omitempty"`
}

func (f *FinancialInstrument21) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument21) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument21) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument21) SetProductGroup(value string) {
	f.ProductGroup = (*Max140Text)(&value)
}

func (f *FinancialInstrument21) SetUmbrellaName(value string) {
	f.UmbrellaName = (*Max35Text)(&value)
}

func (f *FinancialInstrument21) SetBaseCurrency(value string) {
	f.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument21) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument21) SetCountryOfDomicile(value string) {
	f.CountryOfDomicile = (*CountryCode)(&value)
}

func (f *FinancialInstrument21) AddRegisteredDistributionCountry(value string) {
	f.RegisteredDistributionCountry = append(f.RegisteredDistributionCountry, (*CountryCode)(&value))
}
