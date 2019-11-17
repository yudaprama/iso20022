package model

// Information about a party's account.
type AccountParties12 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties7Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation10 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation10 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation10 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation10 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation10 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation10 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty7 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation10 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation10 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties12) AddPrincipalAccountParty() *AccountParties7Choice {
	a.PrincipalAccountParty = new(AccountParties7Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties12) AddSecondaryOwner() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties12) AddBeneficiary() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties12) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties12) AddLegalGuardian() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties12) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties12) AddAdministrator() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties12) AddOtherParty() *ExtendedParty7 {
	newValue := new(ExtendedParty7)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties12) AddGranter() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties12) AddSettlor() *InvestmentAccountOwnershipInformation10 {
	newValue := new(InvestmentAccountOwnershipInformation10)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}
