package model

// Information used to calculate the tax.
type TaxCalculationInformation4 struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain2Code `xml:"EUCptlGn,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUCapitalGain *Extended350Code `xml:"XtndedEUCptlGn,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Percentage of grandfathered debt claim.
	PercentageGrandfatheredDebt *PercentageRate `xml:"PctgGrdfthdDebt,omitempty"`

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatus1Code `xml:"EUDvddSts,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUDividendStatus *Extended350Code `xml:"XtndedEUDvddSts,omitempty"`
}

func (t *TaxCalculationInformation4) SetEUCapitalGain(value string) {
	t.EUCapitalGain = (*EUCapitalGain2Code)(&value)
}

func (t *TaxCalculationInformation4) SetExtendedEUCapitalGain(value string) {
	t.ExtendedEUCapitalGain = (*Extended350Code)(&value)
}

func (t *TaxCalculationInformation4) SetPercentageOfDebtClaim(value string) {
	t.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation4) SetPercentageGrandfatheredDebt(value string) {
	t.PercentageGrandfatheredDebt = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation4) SetTaxableIncomePerDividend(value, currency string) {
	t.TaxableIncomePerDividend = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCalculationInformation4) SetEUDividendStatus(value string) {
	t.EUDividendStatus = (*EUDividendStatus1Code)(&value)
}

func (t *TaxCalculationInformation4) SetExtendedEUDividendStatus(value string) {
	t.ExtendedEUDividendStatus = (*Extended350Code)(&value)
}
