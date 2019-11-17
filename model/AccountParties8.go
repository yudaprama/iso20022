package model

// Party related to an investment account.
type AccountParties8 struct {

	// Specifies the type of modification to be applied on a set of information.
	ModificationScopeIndication *DataModification1Code `xml:"ModScpIndctn"`

	// Main party associated with the account.
	PrincipalAccountParty *AccountParties3Choice `xml:"PrncplAcctPty,omitempty"`

	// Entity that is not the primary owner when the ownership of the investment account is split among several owners.
	SecondaryOwner []*InvestmentAccountOwnershipInformation8 `xml:"ScndryOwnr,omitempty"`

	// Ultimate party that is entitled to either receive the benefits of the ownership of a financial instrument, or to be paid/credited as a result of a transfer.
	Beneficiary []*InvestmentAccountOwnershipInformation8 `xml:"Bnfcry,omitempty"`

	// Entity that was given the authority by another entity to act on its behalf.
	PowerOfAttorney []*InvestmentAccountOwnershipInformation8 `xml:"PwrOfAttny,omitempty"`

	// Entity that has been appointed by a legal authority to act on behalf of a person judged to be incapacitated.
	LegalGuardian []*InvestmentAccountOwnershipInformation8 `xml:"LglGuardn,omitempty"`

	// Deceased's estate, or successor, to whom the respective percentage of ownership will be transferred upon the death of one of the owners.
	SuccessorOnDeath *InvestmentAccountOwnershipInformation8 `xml:"SucssrOnDth,omitempty"`

	// Entity that has been appointed by a legal authorithy to act on behalf of a person or organisation that has gone bankrupt.
	Administrator []*InvestmentAccountOwnershipInformation8 `xml:"Admstr,omitempty"`

	// An other type of party.
	OtherParty []*ExtendedParty5 `xml:"OthrPty,omitempty"`

	// Granter role in the hedge funds industry.
	Granter []*InvestmentAccountOwnershipInformation8 `xml:"Grntr,omitempty"`

	// Settler role in the hedge funds industry.
	Settler []*InvestmentAccountOwnershipInformation8 `xml:"Sttlr,omitempty"`
}

func (a *AccountParties8) SetModificationScopeIndication(value string) {
	a.ModificationScopeIndication = (*DataModification1Code)(&value)
}

func (a *AccountParties8) AddPrincipalAccountParty() *AccountParties3Choice {
	a.PrincipalAccountParty = new(AccountParties3Choice)
	return a.PrincipalAccountParty
}

func (a *AccountParties8) AddSecondaryOwner() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.SecondaryOwner = append(a.SecondaryOwner, newValue)
	return newValue
}

func (a *AccountParties8) AddBeneficiary() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Beneficiary = append(a.Beneficiary, newValue)
	return newValue
}

func (a *AccountParties8) AddPowerOfAttorney() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.PowerOfAttorney = append(a.PowerOfAttorney, newValue)
	return newValue
}

func (a *AccountParties8) AddLegalGuardian() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.LegalGuardian = append(a.LegalGuardian, newValue)
	return newValue
}

func (a *AccountParties8) AddSuccessorOnDeath() *InvestmentAccountOwnershipInformation8 {
	a.SuccessorOnDeath = new(InvestmentAccountOwnershipInformation8)
	return a.SuccessorOnDeath
}

func (a *AccountParties8) AddAdministrator() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Administrator = append(a.Administrator, newValue)
	return newValue
}

func (a *AccountParties8) AddOtherParty() *ExtendedParty5 {
	newValue := new(ExtendedParty5)
	a.OtherParty = append(a.OtherParty, newValue)
	return newValue
}

func (a *AccountParties8) AddGranter() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Granter = append(a.Granter, newValue)
	return newValue
}

func (a *AccountParties8) AddSettler() *InvestmentAccountOwnershipInformation8 {
	newValue := new(InvestmentAccountOwnershipInformation8)
	a.Settler = append(a.Settler, newValue)
	return newValue
}
