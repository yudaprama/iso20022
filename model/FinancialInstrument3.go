package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument3 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification1Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *ActiveOrHistoricCurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument3) AddIdentification() *SecurityIdentification1Choice {
	f.Identification = new(SecurityIdentification1Choice)
	return f.Identification
}

func (f *FinancialInstrument3) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument3) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument3) SetRequestedNAVCurrency(value string) {
	f.RequestedNAVCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument3) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument3) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument3) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}
