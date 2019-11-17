package model

// Tax related to an investment fund order.
type InformativeTax1 struct {

	// Amount included in the dividend that corresponds to gains directly or indirectly derived from interest payment in the scope of the European Directive on taxation of savings income in the form of interest payments.
	TaxableIncomePerDividend *ActiveCurrencyAndAmount `xml:"TaxblIncmPerDvdd,omitempty"`

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain3Choice `xml:"EUCptlGn,omitempty"`

	// Specifies whether dividend is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUDividendStatus *EUDividendStatusType2Choice `xml:"EUDvddSts,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Information related to a specific tax that is provided for information purposes.
	IndividualTax []*Tax32 `xml:"IndvTax,omitempty"`
}

func (i *InformativeTax1) SetTaxableIncomePerDividend(value, currency string) {
	i.TaxableIncomePerDividend = NewActiveCurrencyAndAmount(value, currency)
}

func (i *InformativeTax1) AddEUCapitalGain() *EUCapitalGain3Choice {
	i.EUCapitalGain = new(EUCapitalGain3Choice)
	return i.EUCapitalGain
}

func (i *InformativeTax1) AddEUDividendStatus() *EUDividendStatusType2Choice {
	i.EUDividendStatus = new(EUDividendStatusType2Choice)
	return i.EUDividendStatus
}

func (i *InformativeTax1) SetPercentageOfDebtClaim(value string) {
	i.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (i *InformativeTax1) AddIndividualTax() *Tax32 {
	newValue := new(Tax32)
	i.IndividualTax = append(i.IndividualTax, newValue)
	return newValue
}
