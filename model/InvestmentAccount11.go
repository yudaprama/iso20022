package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount11 struct {

	// Party that legally owns the account.
	OwnerIdentification []*PartyIdentification1Choice `xml:"OwnrId,omitempty"`

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	AccountIdentification *AccountIdentification1 `xml:"AcctId"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	AccountName *Max35Text `xml:"AcctNm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	AccountDesignation *Max35Text `xml:"AcctDsgnt,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	IntermediaryInformation []*Intermediary1 `xml:"IntrmyInf,omitempty"`

	// Form, ie, ownership, of the security, eg, registered or bearer.
	SecuritiesForm *FormOfSecurity1Code `xml:"SctiesForm,omitempty"`

	// Indicates whether a security exists only as an electronic record, ie, there is no physical document representing the security.
	DematerialisedIndicator *YesNoIndicator `xml:"DmtrlsdInd,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Indicates whether the beneficial ownership certification has been sent, certifying that the beneficial owner is eligible to own a specific investment fund or investment fund class.
	BeneficiaryCertificationIndicator *YesNoIndicator `xml:"BnfcryCertfctnInd,omitempty"`

	// Party related to an account that is not the legal account owner, eg, the power of attorney.
	AccountServicer *PartyIdentification1Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount11) AddOwnerIdentification() *PartyIdentification1Choice {
	newValue := new(PartyIdentification1Choice)
	i.OwnerIdentification = append(i.OwnerIdentification, newValue)
	return newValue
}

func (i *InvestmentAccount11) AddAccountIdentification() *AccountIdentification1 {
	i.AccountIdentification = new(AccountIdentification1)
	return i.AccountIdentification
}

func (i *InvestmentAccount11) SetAccountName(value string) {
	i.AccountName = (*Max35Text)(&value)
}

func (i *InvestmentAccount11) SetAccountDesignation(value string) {
	i.AccountDesignation = (*Max35Text)(&value)
}

func (i *InvestmentAccount11) AddIntermediaryInformation() *Intermediary1 {
	newValue := new(Intermediary1)
	i.IntermediaryInformation = append(i.IntermediaryInformation, newValue)
	return newValue
}

func (i *InvestmentAccount11) SetSecuritiesForm(value string) {
	i.SecuritiesForm = (*FormOfSecurity1Code)(&value)
}

func (i *InvestmentAccount11) SetDematerialisedIndicator(value string) {
	i.DematerialisedIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount11) SetIncomePreference(value string) {
	i.IncomePreference = (*IncomePreference1Code)(&value)
}

func (i *InvestmentAccount11) SetBeneficiaryCertificationIndicator(value string) {
	i.BeneficiaryCertificationIndicator = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccount11) AddAccountServicer() *PartyIdentification1Choice {
	i.AccountServicer = new(PartyIdentification1Choice)
	return i.AccountServicer
}
