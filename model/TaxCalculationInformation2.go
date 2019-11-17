package model

// Information used to calculate the tax.
type TaxCalculationInformation2 struct {

	// Specifies whether capital gain is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June), or an income realised upon sale, a refund or redemption of shares and units, etc.
	EUCapitalGain *EUCapitalGain1 `xml:"EUCptlGn,omitempty"`

	// Percentage of the underlying assets of the funds that represents a debt and is in the scope of the European directive on taxation of savings income in the form of interest payments (Council Directive 2003/48/EC 3 June).
	PercentageOfDebtClaim *PercentageRate `xml:"PctgOfDebtClm,omitempty"`

	// Percentage of grandfathered debt claim.
	PercentageGrandfatheredDebt *PercentageRate `xml:"PctgGrdfthdDebt,omitempty"`
}

func (t *TaxCalculationInformation2) AddEUCapitalGain() *EUCapitalGain1 {
	t.EUCapitalGain = new(EUCapitalGain1)
	return t.EUCapitalGain
}

func (t *TaxCalculationInformation2) SetPercentageOfDebtClaim(value string) {
	t.PercentageOfDebtClaim = (*PercentageRate)(&value)
}

func (t *TaxCalculationInformation2) SetPercentageGrandfatheredDebt(value string) {
	t.PercentageGrandfatheredDebt = (*PercentageRate)(&value)
}
