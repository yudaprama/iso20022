package model

// Party related to an investment account.
type AccountParties9 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties4Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation9 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation9 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation9 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation9 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation9 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation9 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty6 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation9 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation9 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties9) AddPrincipalAccountParty() *AccountParties4Choice {
	a.PrincipalAccountParty = new(AccountParties4Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties9) AddSecondaryOwner() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties9) AddBeneficiary() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties9) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties9) AddLegalGuardian() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties9) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties9) AddAdministrator() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties9) AddOtherParty() *ExtendedParty6 {
	newValue := new(ExtendedParty6)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties9) AddGranter() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties9) AddSettler() *InvestmentAccountOwnershipInformation9 {
	newValue := new(InvestmentAccountOwnershipInformation9)
	a.Settler = append(a.Settler, newValue)
	return newValue
}
