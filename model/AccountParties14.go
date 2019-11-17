package model

// Information about a party's account.
type AccountParties14 struct {

	// Specifies the type of modification to be applied.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties9Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner *InvestmentAccountOwnershipInformation13 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary *InvestmentAccountOwnershipInformation13 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney *InvestmentAccountOwnershipInformation13 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian *InvestmentAccountOwnershipInformation13 `xml:"LglGuardn,omitempty"`

	// Entity that holds shares/units on behalf of a legal minor. Although the account is registered under the name of the minor, the custodian retains control of the account.
	CustodianForMinor *InvestmentAccountOwnershipInformation13 `xml:"CtdnForMnr,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation13 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator *InvestmentAccountOwnershipInformation13 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty *ExtendedParty10 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter *InvestmentAccountOwnershipInformation13 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor *InvestmentAccountOwnershipInformation13 `xml:"Sttlr,omitempty"`

	// Party for which shares are to be registered.
	RegisteredShareholderName *RegisteredShareholderName1Choice `xml:"RegdShrhldrNm,omitempty"`
}

func (a *AccountParties14) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties14) AddPrincipalAccountParty() *AccountParties9Choice {
	a.PrincipalAccountParty = new(AccountParties9Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties14) AddSecondaryOwner() *InvestmentAccountOwnershipInformation13 {
	a.SecondaryOwner = new(InvestmentAccountOwnershipInformation13)
	return a.SecondaryOwner
}

func (a *AccountParties14) AddBeneficiary() *InvestmentAccountOwnershipInformation13 {
	a.Beneficiary = new(InvestmentAccountOwnershipInformation13)
	return a.Beneficiary
}

func (a *AccountParties14) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation13 {
	a.PowerOfAttorney = new(InvestmentAccountOwnershipInformation13)
	return a.PowerOfAttorney
}

func (a *AccountParties14) AddLegalGuardian() *InvestmentAccountOwnershipInformation13 {
	a.LegalGuardian = new(InvestmentAccountOwnershipInformation13)
	return a.LegalGuardian
}

func (a *AccountParties14) AddCustodianForMinor() *InvestmentAccountOwnershipInformation13 {
	a.CustodianForMinor = new(InvestmentAccountOwnershipInformation13)
	return a.CustodianForMinor
}

func (a *AccountParties14) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation13 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation13)
	return a.SuccessorOnDeath
}

func (a *AccountParties14) AddAdministrator() *InvestmentAccountOwnershipInformation13 {
	a.Administrator = new(InvestmentAccountOwnershipInformation13)
	return a.Administrator
}

func (a *AccountParties14) AddOtherParty() *ExtendedParty10 {
	a.OtherParty = new(ExtendedParty10)
	return a.OtherParty
}

func (a *AccountParties14) AddGranter() *InvestmentAccountOwnershipInformation13 {
	a.Granter = new(InvestmentAccountOwnershipInformation13)
	return a.Granter
}

func (a *AccountParties14) AddSettlor() *InvestmentAccountOwnershipInformation13 {
	a.Settlor = new(InvestmentAccountOwnershipInformation13)
	return a.Settlor
}

func (a *AccountParties14) AddRegisteredShareholderName() *RegisteredShareholderName1Choice {
	a.RegisteredShareholderName = new(RegisteredShareholderName1Choice)
	return a.RegisteredShareholderName
}
