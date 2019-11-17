package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount21 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Counterparties eligibility as defined by article 24 of the EU MiFID Directive applicable to transactions executed by investment firms for eligible counterparties.
	OrderOriginatorEligibility *OrderOriginatorEligibility1Code `xml:"OrdrOrgtrElgblty,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount21) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount21) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount21) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount21) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount21) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount21) SetOrderOriginatorEligibility(value string) {
	i.OrderOriginatorEligibility = (*OrderOriginatorEligibility1Code)(&value)
}

func (i *InvestmentAccount21) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}
