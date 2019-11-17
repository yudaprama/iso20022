package model

// Distinct pool of financial instruments managed by a single investment policy. May or not be part of an umbrella fund.The pool is issued in at least one investment fund class.
type FundIdentification2 struct {

	// Identification of the investment fund.
	FundIdentification *Max35Text `xml:"FndId"`

	// Identifies the account of the fund held with the custodian.
	AccountIdentificationWithCustodian *Max35Text `xml:"AcctIdWthCtdn,omitempty"`

	// Identification of the custodian which services the account of the fund.
	CustodianIdentification *PartyIdentification8Choice `xml:"CtdnId,omitempty"`
}

func (f *FundIdentification2) SetFundIdentification(value string) {
	f.FundIdentification = (*Max35Text)(&value)
}

func (f *FundIdentification2) SetAccountIdentificationWithCustodian(value string) {
	f.AccountIdentificationWithCustodian = (*Max35Text)(&value)
}

func (f *FundIdentification2) AddCustodianIdentification() *PartyIdentification8Choice {
	f.CustodianIdentification = new(PartyIdentification8Choice)
	return f.CustodianIdentification
}
