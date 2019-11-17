package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument20 struct {

	// Indicate whether or note it is possible to hold bearer units/shares in this class in certified form
	PhysicalBearerSecurities *YesNoIndicator `xml:"PhysBrScties"`

	// Indicate whether or not it is possible to hold bearer units/shares in paperless form
	DematerialisedBearerSecurities *YesNoIndicator `xml:"DmtrlsdBrScties"`

	// Indicate whether or not it is possible to hold registered units/shares in this class in paperless form
	PhysicalRegisteredSecurities *YesNoIndicator `xml:"PhysRegdScties"`

	// Indicate whether or not it is possible to hold registered units/shares in this class in paperless form
	DematerialisedRegisteredSecurities *YesNoIndicator `xml:"DmtrlsdRegdScties"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy"`

	// Dividend policy of the fund, eg, cash, units.
	DividendPolicy *DividendPolicy1Code `xml:"DvddPlcy,omitempty"`

	// Frequency with which the income is allocated to investors.
	DividendFrequency *EventFrequency5Code `xml:"DvddFrqcy,omitempty"`

	// Frequency with which the reinvestment takes place,  This is the same or less than the dividend frequency,
	ReinvestmentFrequency *EventFrequency5Code `xml:"RinvstmtFrqcy,omitempty"`

	// Front end charge on subscription orders for this class can be applied.
	FrontEndLoad *YesNoIndicator `xml:"FrntEndLd"`

	// Exit charge (eg. CDSC) on redemption orders for this class can be applied.
	BackEndLoad *YesNoIndicator `xml:"BckEndLd"`

	// If a separate fee for switching between sub-funds of the same umbrella can be applied
	SwitchFee *YesNoIndicator `xml:"SwtchFee"`

	// Indicates whether the investment fund class is subject to the European Union Saving Directive.
	EUSavingsDirective *EUSavingsDirective1Code `xml:"EUSvgsDrctv"`
}

func (f *FinancialInstrument20) SetPhysicalBearerSecurities(value string) {
	f.PhysicalBearerSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDematerialisedBearerSecurities(value string) {
	f.DematerialisedBearerSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetPhysicalRegisteredSecurities(value string) {
	f.PhysicalRegisteredSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDematerialisedRegisteredSecurities(value string) {
	f.DematerialisedRegisteredSecurities = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument20) SetDividendPolicy(value string) {
	f.DividendPolicy = (*DividendPolicy1Code)(&value)
}

func (f *FinancialInstrument20) SetDividendFrequency(value string) {
	f.DividendFrequency = (*EventFrequency5Code)(&value)
}

func (f *FinancialInstrument20) SetReinvestmentFrequency(value string) {
	f.ReinvestmentFrequency = (*EventFrequency5Code)(&value)
}

func (f *FinancialInstrument20) SetFrontEndLoad(value string) {
	f.FrontEndLoad = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetBackEndLoad(value string) {
	f.BackEndLoad = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetSwitchFee(value string) {
	f.SwitchFee = (*YesNoIndicator)(&value)
}

func (f *FinancialInstrument20) SetEUSavingsDirective(value string) {
	f.EUSavingsDirective = (*EUSavingsDirective1Code)(&value)
}
