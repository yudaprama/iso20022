package model

// Information regarding the total amount of taxes.
type TotalTaxes3 struct {

	// Total value of the taxes for a specific order.
	TotalAmountOfTaxes *ActiveCurrencyAnd13DecimalAmount `xml:"TtlAmtOfTaxs,omitempty"`

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveCurrencyAndAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain2Code `xml:"EUCptlGn,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUCapitalGain *Extended350Code `xml:"XtndedEUCptlGn,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatus1Code `xml:"EUDvddSts,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	ExtendedEUDividendStatus *Extended350Code `xml:"XtndedEUDvddSts,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Information related to a specific tax.
	TaxDetails []*Tax14 `xml:"TaxDtls,omitempty"`
}

func (t *TotalTaxes3) SetTotalAmountOfTaxes(value, currency string) {
	t.TotalAmountOfTaxes = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TotalTaxes3) SetTaxableIncomePerDividend(value, currency string) {
	t.TaxableIncomePerDividend = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TotalTaxes3) SetEUCapitalGain(value string) {
	t.EUCapitalGain = (*EUCapitalGain2Code)(&value)
}

func (t *TotalTaxes3) SetExtendedEUCapitalGain(value string) {
	t.ExtendedEUCapitalGain = (*Extended350Code)(&value)
}

func (t *TotalTaxes3) SetEUDividendStatus(value string) {
	t.EUDividendStatus = (*EUDividendStatus1Code)(&value)
}

func (t *TotalTaxes3) SetExtendedEUDividendStatus(value string) {
	t.ExtendedEUDividendStatus = (*Extended350Code)(&value)
}

func (t *TotalTaxes3) SetPercentageOfDebtClaim(value string) {
	t.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (t *TotalTaxes3) AddTaxDetails() *Tax14 {
	newValue := new(Tax14)
	t.TaxDetails = append(t.TaxDetails, newValue)
	return newValue
}
