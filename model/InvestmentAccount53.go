package model

// Information about a securities account and its characteristics.
type InvestmentAccount53 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *Max35Text `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Identification of the owner of the account.
	OwnerIdentification *OwnerIdentification2Choice `xml:"OwnrId,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification70Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount53) SetAccountIdentification(value string) {
	i.AccountIdentification = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount53) AddOwnerIdentification() *OwnerIdentification2Choice {
	i.OwnerIdentification = new(OwnerIdentification2Choice)
	return i.OwnerIdentification
}

func (i *InvestmentAccount53) AddAccountServicer() *PartyIdentification70Choice {
	i.AccountServicer = new(PartyIdentification70Choice)
	return i.AccountServicer
}
