package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount58 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification113 `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification113 `xml:"AcctSvcr,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Sub-account of the master or omnibus account.
	SubAccountDetails *SubAccount6 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount58) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount58) AddOwnerIdentification() *PartyIdentification113 {
	newValue := new(PartyIdentification113)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount58) AddAccountServicer() *PartyIdentification113 {
	i.AccountServicer = new(PartyIdentification113)
	return i.AccountServicer
}

func (i *InvestmentAccount58) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *InvestmentAccount58) AddSubAccountDetails() *SubAccount6 {
	i.SubAccountDetails = new(SubAccount6)
	return i.SubAccountDetails
}
