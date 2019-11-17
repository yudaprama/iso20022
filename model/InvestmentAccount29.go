package model

// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
type InvestmentAccount29 struct {

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Supplementary registration information applying to a specific block of units for dealing and reporting purposes. The supplementary registration information may be used when all the units are registered, for example, to a funds supermarket, but holdings for each investor have to reconciled individually.
	Designation *Max35Text `xml:"Dsgnt,omitempty"`

	// Legal form of the fund, eg, UCITS, SICAV, OEIC, Unit Trust, and FCP.
	FundType *Max35Text `xml:"FndTp,omitempty"`

	// Name of the investment fund family.
	FundFamilyName *Max350Text `xml:"FndFmlyNm,omitempty"`

	// Detailed information about the investment fund associated to the account.
	SecurityDetails *FinancialInstrument10 `xml:"SctyDtls,omitempty"`

	// Identification of an individual person whom legally owns the account.
	IndividualOwnerIdentification *IndividualPersonIdentificationChoice `xml:"IndvOwnrId,omitempty"`

	// Identification of an organisation that legally owns the account.
	OrganisationOwnerIdentification *PartyIdentification5Choice `xml:"OrgOwnrId,omitempty"`

	// Party that provides services relating to financial products to investors, eg, advice on products and placement of orders for the investment fund.
	Intermediary []*Intermediary7 `xml:"Intrmy,omitempty"`

	// Party that manages the account on behalf of the account owner, that is manages the registration and booking of entries on the account, calculates balances on the account and provides information about the account.
	AccountServicer *PartyIdentification2Choice `xml:"AcctSvcr,omitempty"`
}

func (i *InvestmentAccount29) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetDesignation(value string) {
	i.Designation = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetFundType(value string) {
	i.FundType = (*Max35Text)(&value)
}

func (i *InvestmentAccount29) SetFundFamilyName(value string) {
	i.FundFamilyName = (*Max350Text)(&value)
}

func (i *InvestmentAccount29) AddSecurityDetails() *FinancialInstrument10 {
	i.SecurityDetails = new(FinancialInstrument10)
	return i.SecurityDetails
}

func (i *InvestmentAccount29) AddIndividualOwnerIdentification() *IndividualPersonIdentificationChoice {
	i.IndividualOwnerIdentification = new(IndividualPersonIdentificationChoice)
	return i.IndividualOwnerIdentification
}

func (i *InvestmentAccount29) AddOrganisationOwnerIdentification() *PartyIdentification5Choice {
	i.OrganisationOwnerIdentification = new(PartyIdentification5Choice)
	return i.OrganisationOwnerIdentification
}

func (i *InvestmentAccount29) AddIntermediary() *Intermediary7 {
	newValue := new(Intermediary7)
	i.Intermediary = append(i.Intermediary, newValue)
	return newValue
}

func (i *InvestmentAccount29) AddAccountServicer() *PartyIdentification2Choice {
	i.AccountServicer = new(PartyIdentification2Choice)
	return i.AccountServicer
}
