package model

// Distinct pool of financial instruments managed by a single investment policy. May or may not be part of an umbrella fund.The pool is issued in at least one investment fund class.
type FundIdentification4 struct {

	// Identification of the investment fund.
	FundIdentification *PartyIdentification60 `xml:"FndId"`

	// Identifies the account of the fund held with the custodian.
	AccountIdentificationWithCustodian *Max35Text `xml:"AcctIdWthCtdn,omitempty"`

	// Identification of the custodian which services the account of the fund.
	CustodianIdentification *PartyIdentification73Choice `xml:"CtdnId,omitempty"`
}

func (f *FundIdentification4) AddFundIdentification() *PartyIdentification60 {
	f.FundIdentification = new(PartyIdentification60)
	return f.FundIdentification
}

func (f *FundIdentification4) SetAccountIdentificationWithCustodian(value string) {
	f.AccountIdentificationWithCustodian = (*Max35Text)(&value)
}

func (f *FundIdentification4) AddCustodianIdentification() *PartyIdentification73Choice {
	f.CustodianIdentification = new(PartyIdentification73Choice)
	return f.CustodianIdentification
}
