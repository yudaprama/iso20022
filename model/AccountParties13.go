package model

// Information about a party's account.
type AccountParties13 struct {

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties8Choice `xml:"PrncplAcctPty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation12 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation12 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation12 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation12 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation12 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath []*InvestmentAccountOwnershipInformation12 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation12 `xml:"Admstr,omitempty"`

	// Other type of party.
	OtherParty []*ExtendedParty9 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation12 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation12 `xml:"Sttlr,omitempty"`

	// Party that registers its name with the issuer and the name used for the registration.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties13) AddPrincipalAccountParty() *AccountParties8Choice {
	a.PrincipalAccountParty = new(AccountParties8Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties13) AddSecondaryOwner() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties13) AddBeneficiary() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties13) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties13) AddLegalGuardian() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties13) AddCustodianForMinor() *InvestmentAccountOwnershipInformation12 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation12)
	return a.CustodianForMinor
}

func (a *AccountParties13) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.SuccessorOnDeath = append(a.SuccessorOnDeath, newValue)
	return newValue
}

func (a *AccountParties13) AddAdministrator() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties13) AddOtherParty() *ExtendedParty9 {
	newValue := new(ExtendedParty9)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties13) AddGranter() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties13) AddSettlor() *InvestmentAccountOwnershipInformation12 {
	newValue := new(InvestmentAccountOwnershipInformation12)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}

func (a *AccountParties13) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}
