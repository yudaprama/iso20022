package model

// Party related to an investment account.
type AccountParties6 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties1Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation6 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation6 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation6 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation6 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation6 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation6 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty3 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation6 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation6 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties6) AddPrincipalAccountParty() *AccountParties1Choice {
	a.PrincipalAccountParty = new(AccountParties1Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties6) AddSecondaryOwner() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties6) AddBeneficiary() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties6) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties6) AddLegalGuardian() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties6) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties6) AddAdministrator() *InvestmentAccountOwnershipInformation6 {
	a.Administrator = new(InvestmentAccountOwnershipInformation6)
	return a.Administrator
}

func (a *AccountParties6) AddOtherParty() *ExtendedParty3 {
	newValue := new(ExtendedParty3)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties6) AddGranter() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties6) AddSettler() *InvestmentAccountOwnershipInformation6 {
	newValue := new(InvestmentAccountOwnershipInformation6)
	a.Settler = append(a.Settler, newValue)
	return newValue
}
