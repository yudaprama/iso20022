package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount24 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification2Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId,omitempty"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary10 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`

	// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`
}

func (i *InvestmentAccount24) AddOwnerIdentification() *PartyIdentification2Choice {
	newValue := new(PartyIdentification2Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount24) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount24) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount24) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount24) AddIntermediaryInformation() *Intermediary10 {
	newValue := new(Intermediary10)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount24) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount24) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount24) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount24) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *InvestmentAccount24) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}

func (i *InvestmentAccount24) AddSubAccountDetails() *SubAccount1 {
	i.SubAccountDetails = new(SubAccount1)
	return i.SubAccountDetails
}
