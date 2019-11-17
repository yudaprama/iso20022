package model

// Provides information such as the default fund account identification or the default fund amount.
type DefaultFund1 struct {

	// Specifies the account identification of the clearing member at the ICSD (International Central Securities Depository) or at the central bank.
	DefaultFundAccount *AccountIdentification4Choice `xml:"DfltFndAcct"`

	// Total amount required by the clearing member to participate to the default fund.
	TotalDefaultFundAmount *ActiveCurrencyAndAmount `xml:"TtlDfltFndAmt"`

	// Provides details about the contribution to the default fund by trading venues/products.
	Contribution []*Contribution1 `xml:"Cntrbtn,omitempty"`

	// Additional amount that the clearing member will have to provide to cover a risk increase. This results from a risk management decision depending on central counterparty specific criteria.
	IncreaseCoverageAmount *ActiveCurrencyAndAmount `xml:"IncrCvrgAmt,omitempty"`
}

func (d *DefaultFund1) AddDefaultFundAccount() *AccountIdentification4Choice {
	d.DefaultFundAccount = new(AccountIdentification4Choice)
	return d.DefaultFundAccount
}

func (d *DefaultFund1) SetTotalDefaultFundAmount(value, currency string) {
	d.TotalDefaultFundAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DefaultFund1) AddContribution() *Contribution1 {
	newValue := new(Contribution1)
	d.Contribution = append(d.Contribution, newValue)
	return newValue
}

func (d *DefaultFund1) SetIncreaseCoverageAmount(value, currency string) {
	d.IncreaseCoverageAmount = NewActiveCurrencyAndAmount(value, currency)
}
