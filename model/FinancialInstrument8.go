package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument8 struct {

	// Identification of a security by an ISIN.
	Identification []*SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Currency in which a security is issued or redenominated.
	DenominationCurrency *ActiveOrHistoricCurrencyCode `xml:"DnmtnCcy,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, eg, front end or back end, an income policy, eg, pay out or accumulate, or a trailer policy, eg, with or without. Fund classes are typically denoted by a single character, eg, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, ie, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`

	// Indicates whether the fund has two prices.
	DualFundIndicator *YesNoIndicator `xml:"DualFndInd"`
}

func (f *FinancialInstrument8) AddIdentification() *SecurityIdentification3Choice {
	newValue := new(SecurityIdentification3Choice)
	f.Identification = append(f.Identification, newValue)
	return newValue
}

func (f *FinancialInstrument8) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument8) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument8) SetDenominationCurrency(value string) {
	f.DenominationCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FinancialInstrument8) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument8) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument8) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}

func (f *FinancialInstrument8) SetDualFundIndicator(value string) {
	f.DualFundIndicator = (*YesNoIndicator)(&value)
}
