package model

// Information about a party's account.
type AccountParties11 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties6Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation11 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation11 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation11 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation11 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation11 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation11 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty8 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation11 `xml:"Grntr,omitempty"`

	// Entity that creates a trust or contributes assets to the trust.
	Settlor []*InvestmentAccountOwnershipInformation11 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties11) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties11) AddPrincipalAccountParty() *AccountParties6Choice {
	a.PrincipalAccountParty = new(AccountParties6Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties11) AddSecondaryOwner() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties11) AddBeneficiary() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties11) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties11) AddLegalGuardian() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties11) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation11 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation11)
	return a.SuccessorOnDeath
}

func (a *AccountParties11) AddAdministrator() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties11) AddOtherParty() *ExtendedParty8 {
	newValue := new(ExtendedParty8)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties11) AddGranter() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties11) AddSettlor() *InvestmentAccountOwnershipInformation11 {
	newValue := new(InvestmentAccountOwnershipInformation11)
	a.Settlor = append(a.Settlor, newValue)
	return newValue
}
