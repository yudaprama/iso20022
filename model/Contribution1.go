package model

// Provides additional information such as the contribution account identification or the requirement amount.
type Contribution1 struct {

	// Segregation done by the central counterparty based on trading venues/products or other attributes.
	Account *AccountIdentification4Choice `xml:"Acct,omitempty"`

	// Total contribution required by the clearing member to participate to the default fund.
	RequiredAmount *ActiveCurrencyAndAmount `xml:"ReqrdAmt"`

	// Additional amount that the clearing member will have to provide to cover a risk increase. This results from a risk management decision depending on central counterparty specific criteria.
	IncreaseCoverageAmount *ActiveCurrencyAndAmount `xml:"IncrCvrgAmt,omitempty"`

	// Provides the identification for the non-clearing member.
	NonClearingMember *PartyIdentificationAndAccount31 `xml:"NonClrMmb,omitempty"`
}

func (c *Contribution1) AddAccount() *AccountIdentification4Choice {
	c.Account = new(AccountIdentification4Choice)
	return c.Account
}

func (c *Contribution1) SetRequiredAmount(value, currency string) {
	c.RequiredAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Contribution1) SetIncreaseCoverageAmount(value, currency string) {
	c.IncreaseCoverageAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *Contribution1) AddNonClearingMember() *PartyIdentificationAndAccount31 {
	c.NonClearingMember = new(PartyIdentificationAndAccount31)
	return c.NonClearingMember
}
