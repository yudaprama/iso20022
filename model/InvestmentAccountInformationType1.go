package model

// Select the type(s) of information to be retrieved about a registered investment account.
type InvestmentAccountInformationType1 struct {

	// Indicates if the core investment account information must be selected.
	InvestmentAccount *YesNoIndicator `xml:"InvstmtAcct"`

	// Indicates if the information about account parties must be selected.
	AccountParties *YesNoIndicator `xml:"AcctPties"`

	// Indicates if the information about the intermediaries must be selected.
	Intermediaries *YesNoIndicator `xml:"Intrmies"`

	// Indicates if the information about the investment plan(s) must be selected.
	InvestmentPlan *YesNoIndicator `xml:"InvstmtPlan"`

	// Indicates if the cash settlement information must be selected.
	CashSettlement *YesNoIndicator `xml:"CshSttlm"`

	// Indicates if the Service Level Agreement information must be selected.
	ServiceLevelAgreement *YesNoIndicator `xml:"SvcLvlAgrmt"`
}

func (i *InvestmentAccountInformationType1) SetInvestmentAccount(value string) {
	i.InvestmentAccount = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountInformationType1) SetAccountParties(value string) {
	i.AccountParties = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountInformationType1) SetIntermediaries(value string) {
	i.Intermediaries = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountInformationType1) SetInvestmentPlan(value string) {
	i.InvestmentPlan = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountInformationType1) SetCashSettlement(value string) {
	i.CashSettlement = (*YesNoIndicator)(&value)
}

func (i *InvestmentAccountInformationType1) SetServiceLevelAgreement(value string) {
	i.ServiceLevelAgreement = (*YesNoIndicator)(&value)
}
