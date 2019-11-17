package model

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, for example, dividend option or valuation currency.
type FinancialInstrument49 struct {

	// Identification of a security by an ISIN.
	Identification *SecurityIdentification23Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Financial Instrument Short Name (FISN) expressed in conformance with the ISO 18774 standard.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Additional information about a financial instrument to help identify the instrument.
	SupplementaryIdentification *Max35Text `xml:"SplmtryId,omitempty"`

	// Features of units offered by a fund. For example, a unit may have a specific load structure, for example, front end or back end, an income policy, for example, pay out or accumulate, or a trailer policy, for example, with or without. Fund classes are typically denoted by a single character, for example, 'Class A', 'Class 2'.
	ClassType *Max35Text `xml:"ClssTp,omitempty"`

	// Form, that is, ownership, of the security, for example, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Income policy relating to a class type, that is, if income is paid out or retained in the fund.
	DistributionPolicy *DistributionPolicy1Code `xml:"DstrbtnPlcy,omitempty"`
}

func (f *FinancialInstrument49) AddIdentification() *SecurityIdentification23Choice {
	f.Identification = new(SecurityIdentification23Choice)
	return f.Identification
}

func (f *FinancialInstrument49) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument49) SetShortName(value string) {
	f.ShortName = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetSupplementaryIdentification(value string) {
	f.SupplementaryIdentification = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetClassType(value string) {
	f.ClassType = (*Max35Text)(&value)
}

func (f *FinancialInstrument49) SetSecuritiesForm(value string) {
	f.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (f *FinancialInstrument49) SetDistributionPolicy(value string) {
	f.DistributionPolicy = (*DistributionPolicy1Code)(&value)
}
